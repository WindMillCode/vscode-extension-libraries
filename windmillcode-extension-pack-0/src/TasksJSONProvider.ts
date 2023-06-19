/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as vscode from 'vscode';
import { CreateTaskParams, createTask, letDeveloperKnowAboutAnIssue } from './functions';




export class WMLTasksJSONTaskProvider implements vscode.TaskProvider {
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

class TasksJSONCreateTasksParams extends CreateTaskParams {
  constructor(params:Partial<TasksJSONCreateTasksParams>={}){
    params.taskSource = "tasks"
    super(params)
  }
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    // @ts-ignore
    result = [
      new TasksJSONCreateTasksParams({taskName:"update workspace with latest tasks"}),
    ]
    .map((task)=>{
      return createTask(task)
    })
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    letDeveloperKnowAboutAnIssue(err,'Issue while loading windmillcode TasksJSON tasks.')
  }

	return result;
}
