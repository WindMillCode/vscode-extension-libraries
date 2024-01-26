import os
import shutil
import sys
import re
import logging

# Set up logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def find_dist_info_folders(target_folder):
    """Find all .dist-info directories in the target folder."""
    return [f for f in os.listdir(target_folder) if f.endswith('.dist-info')]

def read_dependencies_from_metadata(metadata_file):
    """Read dependencies from a METADATA file in a .dist-info directory."""
    dependencies = []
    with open(metadata_file, 'r', encoding='utf-8') as file:
        for line in file:
            match = re.match(r'^Requires-Dist: ([^;]+)', line)
            if match:
                dependencies.append(match.group(1).split(' ')[0])
    return dependencies

def build_dependency_graph(target_folder, dist_info_folders):
    """Build a graph of dependencies for all packages."""
    dependency_graph = {}
    for folder in dist_info_folders:
        package_name = folder.split('.')[0].replace('_', '-')
        metadata_file = os.path.join(target_folder, folder, 'METADATA')
        dependencies = read_dependencies_from_metadata(metadata_file)
        for dep in dependencies:
            if dep not in dependency_graph:
                dependency_graph[dep] = set()
            dependency_graph[dep].add(package_name)
    return dependency_graph

def can_remove_package(package_name, dependency_graph, packages_to_remove):
    """Check if a package can be safely removed."""
    if package_name not in dependency_graph:
        return True
    shared_dependencies = [dep for dep in dependency_graph[package_name] if dep not in packages_to_remove]
    if shared_dependencies:
        logging.info(f"Package '{package_name}' has shared dependencies {shared_dependencies} and will not be removed.")
        return False
    return True

def delete_package(package_name, target_folder, dependency_graph, packages_to_remove):
    """Delete package and its unique dependencies."""
    package_files = find_dist_info_folders(target_folder)
    for package_file in package_files:
        if package_file.startswith(package_name.replace('-', '_')):
            metadata_file = os.path.join(target_folder, package_file, 'METADATA')
            if os.path.exists(metadata_file):
                dependencies = read_dependencies_from_metadata(metadata_file)
                for dep in dependencies:
                    if can_remove_package(dep, dependency_graph, packages_to_remove):
                        delete_package(dep, target_folder, dependency_graph, packages_to_remove)
            path = os.path.join(target_folder, package_file)
            shutil.rmtree(path)
            logging.info(f"Removed: {path}")

def main():
    if len(sys.argv) < 3:
        logging.error("Usage: python remove_local_packages.py [package_names] [target_folder]")
        sys.exit(1)

    package_names = sys.argv[1:-1]
    target_folder = sys.argv[-1]

    dist_info_folders = find_dist_info_folders(target_folder)
    dependency_graph = build_dependency_graph(target_folder, dist_info_folders)
    packages_to_remove = set(package_names)

    for package_name in package_names:
        if can_remove_package(package_name, dependency_graph, packages_to_remove):
            delete_package(package_name, target_folder, dependency_graph, packages_to_remove)
        else:
            logging.info(f"Kept: {package_name}")

if __name__ == "__main__":
    main()
