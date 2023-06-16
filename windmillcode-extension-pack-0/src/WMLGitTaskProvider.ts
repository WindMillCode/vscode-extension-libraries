/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as vscode from 'vscode';
import { CreateTaskParams, createTask, getFileFromExtensionDirectory, getOutputChannel, letDeveloperKnowAboutAnIssue } from './functions';




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

class GitCreateTasksParams extends CreateTaskParams {
  constructor(params:Partial<GitCreateTasksParams>={}){
    params.taskSource ="git"
    super(params)
  }
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    // @ts-ignore
    result = [
      new GitCreateTasksParams({taskName:"pushing work to git remote"}),
      new GitCreateTasksParams({taskName:"create branch after merged changes"}),
      new GitCreateTasksParams({taskName:"removing a file from being tracked by git"}),
      new GitCreateTasksParams({taskName:"merge changes from current dev branch"}),
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
