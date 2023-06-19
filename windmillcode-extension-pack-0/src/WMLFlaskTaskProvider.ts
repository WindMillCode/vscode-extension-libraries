/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as vscode from 'vscode';
import { CreateTaskParams, createTask, letDeveloperKnowAboutAnIssue } from './functions';




export class WMLFlaskTaskProvider implements vscode.TaskProvider {
	static WindmillType = 'windmillcode';

	constructor(workspaceRoot: string) {

	}

	public provideTasks(): Thenable<vscode.Task[]> | undefined {

		return getTasks();
	}

	public resolveTask(_task: vscode.Task): vscode.Task | undefined {
    const task = _task.definition.task;
    // A Rake task consists of a task and an optional file as specified in RakeTaskDefinition
    // Make sure that this looks like a Rake task by checking that there is a task.
    if (task) {
      // resolveTask requires that the same definition object be used.
      const definition: vscode.TaskDefinition = <any>_task.definition;
      return new vscode.Task(
        definition,
        _task.scope ?? vscode.TaskScope.Workspace,
        definition.task,
        _task.name,
        _task.execution
      );
    }
    return undefined;
	}
}

class FlaskCreateTasksParams extends CreateTaskParams {
  constructor(params:Partial<FlaskCreateTasksParams>={}){
    params.taskSource = "flask backend"
    super(params)
  }
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    // @ts-ignore
    result = [
      new FlaskCreateTasksParams({taskName:"run"}),
      new FlaskCreateTasksParams({taskName:"test"}),
    ]
    .map((task)=>{
      return createTask(task)
    })
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    letDeveloperKnowAboutAnIssue(err,'Issue while loading windmillcode Flask tasks.')
  }

	return result;
}
