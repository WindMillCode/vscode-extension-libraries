import { mkdirSync } from "fs";
import * as os from "os";
import * as path from "path"
import {  letDeveloperKnowAboutAnIssue } from './functions';
import * as https from 'https';
import * as fs from 'fs';
import * as zlib from 'zlib';
import { exec } from 'child_process';
import { promisify } from 'util';
import * as vscode from 'vscode';
const tar = require("tar")
const semver = require('semver');

let AdmZip = require("adm-zip");


let downloadFile =async (url: string, destinationPath: string): Promise<void> => {
  return new Promise((res,rej)=>{


    let file = fs.createWriteStream(destinationPath);

    https.get(url, (response) => {
      if (response.statusCode !== 200) {
        letDeveloperKnowAboutAnIssue(null,`Error: Failed to download the file. Status Code:${response.statusCode}`);
        return;
      }

      let downloadedBytes = 0;
      let totalBytes = parseInt(response.headers['content-length'] || '0', 10);
      letDeveloperKnowAboutAnIssue(null,`Downloading Go`);
      response.on('data', (chunk) => {
        downloadedBytes += chunk.length;
        let progress = (downloadedBytes / totalBytes) * 100;
        // letDeveloperKnowAboutAnIssue(null,`Downloading... ${progress.toFixed(2)}%`);
      });


      response.pipe(file);

      file.on('finish', () => {
        file.close(() => {
          letDeveloperKnowAboutAnIssue(null,'File download completed successfully.');
          res()
        });
      });
    }).on('error', (err) => {
      letDeveloperKnowAboutAnIssue('Error: Failed to download the file:', err.message);
      rej()
    });
  })
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

let checkGoInstalledInExtension =  (installDir:string,desiredVersion="1.20.6",addedToPath=false)=>{
  return new Promise(async(resolve,rej)=>{
    let executable = "windmillcode_go"
    exec(`${executable} version`,async  (error, stdout, stderr) => {
      if (error) {
        if(!addedToPath){
          letDeveloperKnowAboutAnIssue(null,"Adding executable to the path");
          await addToPath(path.normalize(`${installDir}/bin`))
          setTimeout(()=>{},1500)
          let value = await checkGoInstalledInExtension(installDir,desiredVersion,true)
          return resolve(value)

        }
        letDeveloperKnowAboutAnIssue(error.stack)
        letDeveloperKnowAboutAnIssue("It seems the correct version is not installed on the system or the extension, installing go in the extension location");
        resolve(false);
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        letDeveloperKnowAboutAnIssue(installedVersion)
        letDeveloperKnowAboutAnIssue(desiredVersion)
        if(!semver.gt(desiredVersion,installedVersion)){
          letDeveloperKnowAboutAnIssue(null,`Go is installed and the correct version is in the extension ${stdout.trim()}`);
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


async function addToPath(directory:string) {
  let platform:Partial<NodeJS.Platform> = os.platform()
  // @ts-ignore
  let actions ={
    "win32":{
      "env_setter":"setx PATH "
    },
    "darwin":{
      "env_setter":"export PATH="
    },
    "linux":{
      "env_setter":"export PATH="
    }
  }[platform]
  if(process.env.PATH?.includes(directory)){
    letDeveloperKnowAboutAnIssue(null,"go executable is already on the path")
    return Promise.resolve(true)
  }
  process.env.PATH = `${directory}${path.delimiter}${process.env.PATH}`;
  return new Promise((res,rej)=>{
    exec(`${actions?.env_setter}"${process.env.PATH}"`,(err,stdout,stderr)=>{
      if(err){
        letDeveloperKnowAboutAnIssue(null,err.stack)
        res(false)
      }
      letDeveloperKnowAboutAnIssue(null,"added to path sucessfully")
      res(true)
    })
  })

}


async function copyFile(sourcePath:string, destinationPath:string) {
  const readFileAsync = promisify(fs.readFile);
  const writeFileAsync = promisify(fs.writeFile);

  try {
    const data = await readFileAsync(sourcePath);
    await writeFileAsync(destinationPath, data);
    console.log('File copied successfully.');
  } catch (error) {
    console.error('Error copying the file:', error);
  }
}



export let installGo = async (extensionRoot:string,goVersion="1.20.6",) => {
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
  let goURL = `https://golang.org/dl/go${goVersion}.${platform}-amd64.${goArchiveExt}`;
  if(platform === "win32"){
    goURL = `https://dl.google.com/go/go${goVersion}.windows-amd64.${goArchiveExt}`
  }
  let goArchivePath = path.normalize(
    `${installLocation}/go${goVersion}.${platform}-amd64.${goArchiveExt}`
  );
  let goInstallDir = path.normalize( `${installLocation}/go`);
  letDeveloperKnowAboutAnIssue(null,`Installation Dir ${goInstallDir}`)

  // @ts-ignore
  let executable:boolean |"go" |"windmillcode_go" = await checkGoInstalledInExtension(goInstallDir)
  letDeveloperKnowAboutAnIssue("got here before finishing")
  if(executable === false){
    vscode.window.showInformationMessage("Installing go please wait")

    await downloadFile(goURL,goArchivePath)
    if(platform === "win32"){
      unzipZipFile(goArchivePath,installLocation)
    }
    else{
      unzipTarGz(goArchivePath,installLocation)
    }
    removeFile(goArchivePath)
    if(platform === "win32"){
      await copyFile(
        path.normalize(`${goInstallDir}/bin/${goBinary}`),
        path.normalize(`${goInstallDir}/bin/windmillcode_go.exe`)
      )
    }
    else{
      await copyFile(
        path.normalize(`${goInstallDir}/bin/${goBinary}`),
        path.normalize(`${goInstallDir}/bin/windmillcode_go`)
      )
    }
    await addToPath(path.normalize(`${goInstallDir}/bin/`))
    return  {
      executable:"windmillcode_go",
      alreadyInstalled:false
    }
  }
  else{
    if(executable === "windmillcode_go"){
      await addToPath(path.normalize(`${goInstallDir}/bin/`))
    }
    return  {
      executable,
      alreadyInstalled:true
    }
  }




}
