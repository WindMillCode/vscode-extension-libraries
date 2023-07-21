import * as vscode from 'vscode';
import { OperatingSystem } from './models';
import * as path from 'path'
import * as os from 'os'
const { v4: uuidv4 } = require('uuid');
let _channel: vscode.OutputChannel;
export function getOutputChannel(): vscode.OutputChannel {
	if (!_channel) {
		_channel = vscode.window.createOutputChannel('Windmillcode');
	}
	return _channel;
}

export let letDeveloperKnowAboutAnIssue = (err?:any,msg?:string)=>{
  let channel = getOutputChannel();
  if (err?.stderr) {
    channel.appendLine(err.stderr);
    // vscode.window.showInformationMessage(err.stderr)
    console.log(err.stderr)
  }
  if (err?.stdout) {
    channel.appendLine(err.stdout);
    // vscode.window.showInformationMessage(err.stdout)
    console.log(err.stdout)
  }
  if(msg){
    channel.appendLine(msg);
    // vscode.window.showInformationMessage(msg)
    console.log(msg)
  }
  // channel.show(true);

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

export let createTask = (params = new CreateTaskParams())=>{
  let {shellExecOptions,kind,taskScope,taskName,taskSource,currentOS,executable} = params
  let shellExecutionArgs:string[] = [params.workspaceFolder]
  kind.taskScope = taskScope
  kind.taskName = taskName
  let task = new vscode.Task(
    kind,
    taskScope,
    taskName,
    taskSource,
  );

  task.presentationOptions ={
    focus:true,
    clear:true
  }

  try{
    if(currentOS === OperatingSystem.WINDOWS){
        executable = getFileFromExtensionDirectory(executable,currentOS)

        task.execution = new vscode.ShellExecution(
          executable + " "+shellExecutionArgs[0],
          // shellExecutionArgs,
          shellExecOptions
        )
    }
  }
  catch(e){
    letDeveloperKnowAboutAnIssue(e,'Issue while loading windmillcode' +taskSource +' tasks.')
    return null
  }


  return task

}

export class CreateTaskParams {
  constructor(params:Partial<CreateTaskParams>={}){
    Object.assign(
      this,
      {
        ...params
      }
    )
    if(!this.executable){
      this.executable = this.taskSource + "_" + this.taskName
      this.executable = this.executable.replace(/\s/g, "_")
      // @ts-ignore
      let ext:any = {
        [OperatingSystem.WINDOWS]:".ps1",
        [OperatingSystem.LINUX]:".sh",
        [OperatingSystem.MACOS]:".sh",
      }[this.currentOS]
      this.executable+=ext
    }
  }

  kind: vscode.TaskDefinition = {
    type: 'windmillcode',
    uid:uuidv4()
  }
  currentOS:NodeJS.Platform = os.platform()
  shellExecOptions:vscode.ShellExecutionOptions ={cwd:"."}
  workspaceFolder:string = vscode.workspace.workspaceFolders![0].uri.fsPath
  executable!:string
  taskName: vscode.Task["name"] = "Give me a name!!!!"
  taskSource: vscode.Task["source"] = "Give me a source"
  taskScope=vscode.TaskScope.Workspace
}
