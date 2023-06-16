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
import { CreateTaskParams, createTask, getFileFromExtensionDirectory, getOutputChannel, letDeveloperKnowAboutAnIssue } from './functions';
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

class GitCreakTasksParams extends CreateTaskParams {
  constructor(params:Partial<GitCreakTasksParams>={}){
    super(params)
  }
  override taskSource ="git";
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    // @ts-ignore
    result = [
      new GitCreakTasksParams({taskName:"pushing work to git remote"}),
      new GitCreakTasksParams({taskName:"create branch after merged changes"}),
      new GitCreakTasksParams({taskName:"removing a file from being tracked by git"}),

    ]
    .map((task)=>{
      return createTask(task)
    })
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    letDeveloperKnowAboutAnIssue(err,'Issue while loading windmillcode git tasks.')
  }

	return result;
}
