import * as vscode from 'vscode';
import { OperatingSystem } from './models';
import * as path from 'path'
let _channel: vscode.OutputChannel;
export function getOutputChannel(): vscode.OutputChannel {
	if (!_channel) {
		_channel = vscode.window.createOutputChannel('Console');
	}
	return _channel;
}

export let letDeveloperKnowAboutAnIssue = (err:any,msg:string)=>{
  let channel = getOutputChannel();
  if (err.stderr) {
    channel.appendLine(err.stderr);
  }
  if (err.stdout) {
    channel.appendLine(err.stdout);
  }
  channel.appendLine(msg);
  channel.show(true);
}


export let getFileFromExtensionDirectory =(relativeFilePath:string,currentOS:NodeJS.Platform)=>{
  let myExtension = vscode.extensions.getExtension('windmillcode-publisher-0.windmillcode-extension-pack-0');

  if(myExtension){

    return path.join(
      myExtension.extensionPath,
      'task_files',
       currentOS,
       relativeFilePath);
  }
  throw new Error("File not found")

}
