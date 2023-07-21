import { mkdirSync } from "fs";
import * as os from "os";
import * as path from "path"
import {  letDeveloperKnowAboutAnIssue } from './functions';
import * as https from 'https';
import * as fs from 'fs';
import * as zlib from 'zlib';
import { exec } from 'child_process';
const tar = require("tar")
const semver = require('semver');

let AdmZip = require("adm-zip");


let downloadFile = (url: string, destinationPath: string): void => {
  let file = fs.createWriteStream(destinationPath);

  https.get(url, (response) => {
    if (response.statusCode !== 200) {
      letDeveloperKnowAboutAnIssue(null,`Error: Failed to download the file. Status Code:${response.statusCode}`);
      return;
    }

    let downloadedBytes = 0;
    let totalBytes = parseInt(response.headers['content-length'] || '0', 10);

    response.on('data', (chunk) => {
      downloadedBytes += chunk.length;
      let progress = (downloadedBytes / totalBytes) * 100;
      letDeveloperKnowAboutAnIssue(null,`Downloading... ${progress.toFixed(2)}%`);
    });


    response.pipe(file);

    file.on('finish', () => {
      file.close(() => {
        letDeveloperKnowAboutAnIssue(null,'File download completed successfully.');
      });
    });
  }).on('error', (err) => {
    letDeveloperKnowAboutAnIssue('Error: Failed to download the file:', err.message);
  });
};



let unzipZipFile = (zipFilePath: string, extractToPath: string): void => {
  let zip = new AdmZip(zipFilePath);

  try {
    zip.extractAllTo(extractToPath, true);
    letDeveloperKnowAboutAnIssue('File successfully unzipped!');
  } catch (err:any) {
    letDeveloperKnowAboutAnIssue(`Error unzipping the file:${err.message}`);
  }
};

let unzipTarGz = (tarGzFilePath: string, destinationFolder: string): void => {
  const tarGzReadStream = fs.createReadStream(tarGzFilePath);
  const tarExtractStream = tar.extract({ cwd: destinationFolder });

  tarGzReadStream.pipe(zlib.createGunzip()).pipe(tarExtractStream);

  tarExtractStream.on('finish', () => {
    letDeveloperKnowAboutAnIssue(null,`Successfully extracted files to ${destinationFolder}`);
  });

  tarExtractStream.on('error', (err:any) => {
    letDeveloperKnowAboutAnIssue(err);
  });
};

let removeFile = (filePath: string): void => {
  fs.unlink(filePath, (err) => {
    if (err) {
      letDeveloperKnowAboutAnIssue(`Error while removing the file:${JSON.stringify(err,null,4)}`);
    } else {
      letDeveloperKnowAboutAnIssue(null,`File ${filePath} has been removed successfully.`);
    }
  });
};

let checkGoInstalledInExtension =(installDir:string,desiredVersion:string)=>{
  return new Promise((resolve,rej)=>{
    let executable = path.normalize(`${installDir}/bin/go`)
    exec(`${executable} version`, (error, stdout, stderr) => {
      if (error) {
        letDeveloperKnowAboutAnIssue(null,"It seems the correct version is not install on the system or the extension, installing go in the extension location");
        resolve(false);
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        letDeveloperKnowAboutAnIssue(installedVersion)
        letDeveloperKnowAboutAnIssue(desiredVersion)
        if(!semver.gt(desiredVersion,installedVersion)){
          letDeveloperKnowAboutAnIssue(null,`Go is installed and the correct version is in the extension.${stdout.trim()}`);
          resolve(executable);
        }
        else{
          letDeveloperKnowAboutAnIssue(null,`It seems the correct version is not install on the system or the extension installing go on the extension.${stdout.trim()}`);
          resolve(false)
        }
      }
    });
  })
}

const checkGoInstalled = (installDir:string,desiredVersion="1.20.6") => {
  return new Promise((resolve, reject) => {
    exec('go version', (error, stdout, stderr) => {
      if (error) {
        letDeveloperKnowAboutAnIssue(null,"It looks as if go is not installed on the computer becuase it does not seem to be on the PATH or a path this program can access. Checking if its installed in the extension location");
        resolve(checkGoInstalledInExtension(installDir,desiredVersion))
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        if(!semver.gt(desiredVersion,installedVersion)){
          letDeveloperKnowAboutAnIssue(null,`Go is installed and the correct version is on the system.${stdout.trim()}`);
          resolve("go");
        }
        else{
          letDeveloperKnowAboutAnIssue(null,"System Go is installed but it not the needed version, checking the extension for Go");
          resolve(checkGoInstalledInExtension(installDir,desiredVersion))
        }
      }
    });
  });
};




export let installGo = async (extensionRoot:string,goVersion="1.20.5",) => {
  // Change these values as needed
  let installLocation = path.normalize(extensionRoot+"/task_files")
  letDeveloperKnowAboutAnIssue(installLocation)

  try {
    mkdirSync(installLocation, { recursive: true });
  } catch (err) {
    letDeveloperKnowAboutAnIssue("Error creating installation directory:");
    process.exit(1);
  }

  // Determine the platform and letruct platform-specific commands and paths
  let platform = os.platform();
  let goBinary, goArchiveExt, extractCommand;
  if (platform === "win32") {
    goBinary = "go.exe";
    goArchiveExt = "zip";
  } else if (platform === "darwin") {
    goBinary = "go";
    goArchiveExt = "tar.gz";
  } else if (platform === "linux") {
    goBinary = "go";
    goArchiveExt = "tar.gz";
  } else {
    letDeveloperKnowAboutAnIssue("Unsupported platform:", platform);
    process.exit(1);
  }

  // Download and install Go
  // letDeveloperKnowAboutAnIssue(null,`Downloading Go version ${goVersion} for ${platform}...`);
  let goURL = `https://golang.org/dl/go${goVersion}.${platform}-amd64.${goArchiveExt}`;
  if(platform === "win32"){
    goURL = `https://dl.google.com/go/go${goVersion}.windows-amd64.${goArchiveExt}`
  }
  let goArchivePath = path.normalize(
    `${installLocation}/go${goVersion}.${platform}-amd64.${goArchiveExt}`
  );
  let goInstallDir = path.normalize( `${installLocation}/go`);
  letDeveloperKnowAboutAnIssue(null,goInstallDir)
  let executable:any = await checkGoInstalled(goInstallDir)
  letDeveloperKnowAboutAnIssue(null,executable)
  if(executable === false){
    downloadFile(goURL,goArchivePath)
    if(platform === "win32"){
      unzipZipFile(goArchivePath,installLocation)
    }
    else{
      unzipTarGz(goArchivePath,installLocation)
    }
    removeFile(goArchivePath)
    return   path.normalize(`${goInstallDir}/bin/go`)
  }
  else{
    return executable
  }




}
