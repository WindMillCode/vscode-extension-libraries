/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import * as path from 'path';
import * as fs from 'fs';
import * as cp from 'child_process';
import * as vscode from 'vscode';
import { error } from 'console';
import { OperatingSystem, WMLTaskDefinition } from './models';
import { getFileFromExtensionDirectory, getOutputChannel, letDeveloperKnowAboutAnIssue } from './functions';
import * as os from 'os'



export class WMLGitTaskProvider implements vscode.TaskProvider {
	static WindmillType = 'windmillcode';

	constructor(workspaceRoot: string) {

	}

	public provideTasks(): Thenable<vscode.Task[]> | undefined {

		return getTasks();
	}

	public resolveTask(_task: vscode.Task): vscode.Task | undefined {
    return undefined
	}
}




let pushToGitRemote = (kind: vscode.TaskDefinition,currentOS:NodeJS.Platform,workspaceFolder:string)=>{
  let executable:string
  let shellExecutionArgs:string[] = [workspaceFolder]
  let shellExecOptions:vscode.ShellExecutionOptions ={cwd:"."}
  let task = new vscode.Task(
    kind,
    vscode.TaskScope.Workspace,
    "pushing work to git remote",
    'git',
  );

  try{
    if(currentOS === OperatingSystem.WINDOWS){
        executable = getFileFromExtensionDirectory("pushing_work_to_git_remote.ps1",currentOS)

        task.execution = new vscode.ShellExecution(
          executable + " "+shellExecutionArgs[0],
          // shellExecutionArgs,
          shellExecOptions
        )
    }
  }
  catch(e){
    letDeveloperKnowAboutAnIssue(e,'Issue while loading windmillcode git tasks.')
    return null
  }


  return task
}

let createBranchAfterMergedChanges = (kind: vscode.TaskDefinition,currentOS:NodeJS.Platform,workspaceFolder:string)=>{
  let executable:string
  let shellExecutionArgs:string[] = [workspaceFolder]
  let shellExecOptions:vscode.ShellExecutionOptions ={cwd:"."}
  let task = new vscode.Task(
    kind,
    vscode.TaskScope.Workspace,
    "create branch after merged changes",
    'git',
  );

  try{
    if(currentOS === OperatingSystem.WINDOWS){
        executable = getFileFromExtensionDirectory("create_branch_after_merged_changes.ps1",currentOS)

        task.execution = new vscode.ShellExecution(
          executable + " "+shellExecutionArgs[0],
          // shellExecutionArgs,
          shellExecOptions
        )
    }
  }
  catch(e){
    letDeveloperKnowAboutAnIssue(e,'Issue while loading windmillcode git tasks.')
    return null
  }


  return task
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {
    const kind: vscode.TaskDefinition = {
      type: 'windmillcode'
    };
    let currentOS = os.platform();
    let workspaceFolder = vscode.workspace.workspaceFolders![0].uri.fsPath
    // @ts-ignore
    result = [
      pushToGitRemote(kind,currentOS,workspaceFolder),
      createBranchAfterMergedChanges(kind,currentOS,workspaceFolder)
    ]
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    letDeveloperKnowAboutAnIssue(err,'Issue while loading windmillcode git tasks.')
  }

	return result;
}
