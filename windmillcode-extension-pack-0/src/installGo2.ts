import { mkdirSync } from "fs";
import * as os from "os";
import * as path from "path"
import {  notifyDeveloper } from './functions';
import * as https from 'https';
import * as fs from 'fs';
import { exec } from 'child_process';
import { promisify } from 'util';
import * as vscode from 'vscode';
const semver = require('semver');
let AdmZip = require("adm-zip");

const fetch = require('node-fetch');
const targz = require('tar.gz2');

let downloadFile =async (url: string, destinationPath: string): Promise<void> => {

  return new Promise((res,rej)=>{
    let file = fs.createWriteStream(destinationPath);
    notifyDeveloper(null,`pulling file ${url}`)
    https.get(url, (response) => {
      // @ts-ignore
      if (![200,301].includes(response.statusCode)) {
        notifyDeveloper(null,`Error: Failed to download the file. Status Code:${response.statusCode}`);
        return;
      }

      let downloadedBytes = 0;
      let totalBytes = parseInt(response.headers['content-length'] || '0', 10);
      notifyDeveloper(null,`Downloading Go`);
      response.on('data', (chunk) => {
        downloadedBytes += chunk.length;
        let progress = (downloadedBytes / totalBytes) * 100;
        // notifyDeveloper(null,`Downloading... ${progress.toFixed(2)}%`);
      });


      response.pipe(file);

      file.on('finish', () => {
        file.close(() => {
          notifyDeveloper(null,'File download completed successfully.');
          res()
        });
      });
    })
    .on('error', (err) => {
      notifyDeveloper('Error: Failed to download the file:', err.message);
      rej()
    });
  })
};


let unzipZipFile = (zipFilePath: string, extractToPath: string): void => {
  let zip = new AdmZip(zipFilePath);

  try {
    zip.extractAllTo(extractToPath, true);
    notifyDeveloper('File successfully unzipped!');
  } catch (err:any) {
    notifyDeveloper(`Error unzipping the file:${err.message}`);
  }
};

let unzipTarGz = async (tarGzFilePath: string, destinationFolder: string): Promise<any> => {
  return targz().extract(tarGzFilePath, destinationFolder)
  .then(function(){
    notifyDeveloper(null,'Job done!');
  })
  //@ts-ignore
  .catch(function(err){
    notifyDeveloper(null,`Something is wrong ${err}`);
  });
};

let removeFile = (filePath: string): void => {
  fs.unlink(filePath, (err) => {
    if (err) {
      notifyDeveloper(`Error while removing the file:${JSON.stringify(err,null,4)}`);
    } else {
      notifyDeveloper(null,`File ${filePath} has been removed successfully.`);
    }
  });
};

let checkGoInstalledInExtension =  (installDir:string,desiredVersion="1.20.7",addedToPath=false)=>{
  return new Promise(async(resolve,rej)=>{
    let executable = "windmillcode_go"
    exec(`${executable} version`,async  (error, stdout, stderr) => {
      if (error) {
        if(!addedToPath){
          notifyDeveloper(null,"Adding executable to the path");
          await addToPath(path.normalize(`${installDir}/bin`))
          setTimeout(()=>{},1500)
          let value = await checkGoInstalledInExtension(installDir,desiredVersion,true)
          return resolve(value)

        }
        notifyDeveloper(error.stack)
        notifyDeveloper("It seems the correct version is not installed on the system or the extension, installing go in the extension location");
        resolve(false);
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        notifyDeveloper(installedVersion)
        notifyDeveloper(desiredVersion)
        if(!semver.gt(desiredVersion,installedVersion)){
          notifyDeveloper(null,`Go is installed and the correct version is in the extension ${stdout.trim()}`);
          resolve(executable);
        }
        else{
          notifyDeveloper(null,`It seems the correct version is not install on the system or the extension installing go on the extension.${stdout.trim()}`);
          resolve(false)
        }
      }
    });
  })
}

const checkGoInstalled = (installDir:string,desiredVersion="1.20.7") => {
  return new Promise((resolve, reject) => {
    exec('go version', (error, stdout, stderr) => {
      if (error) {
        notifyDeveloper(null,"It looks as if go is not installed on the computer becuase it does not seem to be on the PATH or a path this program can access. Checking if its installed in the extension location");
        resolve(checkGoInstalledInExtension(installDir,desiredVersion))
      } else {
        let versionMatch = stdout.match(/go(\d+\.\d+\.\d+)/);
        let installedVersion = versionMatch ? versionMatch[1] : null;
        if(!semver.gt(desiredVersion,installedVersion)){
          notifyDeveloper(null,`Go is installed and the correct version is on the system.${stdout.trim()}`);
          resolve("go");
        }
        else{
          notifyDeveloper(null,"System Go is installed but it not the needed version, checking the extension for Go");
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
    notifyDeveloper(null,"go executable is already on the path")
    return Promise.resolve(true)
  }
  process.env.PATH = `${directory}${path.delimiter}${process.env.PATH}`;
  return new Promise((res,rej)=>{
    exec(`${actions?.env_setter}"${process.env.PATH}"`,(err,stdout,stderr)=>{
      if(err){
        notifyDeveloper(null,err.stack)
        res(false)
      }
      notifyDeveloper(null,"added to path sucessfully")
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



export let installGo = async (extensionRoot:string,goVersion="1.20.7",) => {
  // Change these values as needed
  let installLocation = path.normalize(extensionRoot+"/task_files")
  notifyDeveloper(installLocation)

  try {
    mkdirSync(installLocation, { recursive: true });
  } catch (err) {
    notifyDeveloper("Error creating installation directory:");
    process.exit(1);
  }

  // Determine the platform and letruct platform-specific commands and paths
  let platform = os.platform();
  let goBinary, goArchiveExt
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
    notifyDeveloper(null,`Unsupported platform: ${platform}`, );
    process.exit(1);
  }
  let unzipFn = {
    "tar.gz":unzipTarGz,
    "zip":unzipZipFile
  }[goArchiveExt]

  // Download and install Go



  let goURL = `https://dl.google.com/go/go${goVersion}.${platform}-amd64.${goArchiveExt}`
  if(platform === "win32"){
    goURL = `https://dl.google.com/go/go${goVersion}.windows-amd64.${goArchiveExt}`
  }
  let goArchivePath = path.normalize(
    `${installLocation}/go${goVersion}.${platform}-amd64.${goArchiveExt}`
  );
  let goInstallDir = path.normalize( `${installLocation}/go`);
  notifyDeveloper(null,`Installation Dir ${goInstallDir}`)

  // @ts-ignore
  let executable:boolean |"go" |"windmillcode_go" = await checkGoInstalledInExtension(goInstallDir)


  notifyDeveloper(null, ` Executable ${executable}`)
  if(executable === false){
    vscode.window.showInformationMessage("Installing go please wait")

    await downloadFile(goURL,goArchivePath)
    await unzipFn(goArchivePath,installLocation)


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

    await downloadFile(goURL,goArchivePath)
    let result = await unzipTarGz(goArchivePath,installLocation)
    notifyDeveloper(null,` Download result is ${result}`)
    if(executable === "windmillcode_go"){
      await addToPath(path.normalize(`${goInstallDir}/bin/`))
    }
    return  {
      executable,
      alreadyInstalled:true
    }
  }




}