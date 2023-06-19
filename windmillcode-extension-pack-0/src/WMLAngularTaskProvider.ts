/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as vscode from 'vscode';
import { CreateTaskParams, createTask, getFileFromExtensionDirectory, getOutputChannel, letDeveloperKnowAboutAnIssue } from './functions';




export class WMLAngularTaskProvider implements vscode.TaskProvider {
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

class AngularCreateTasksParams extends CreateTaskParams {
  constructor(params:Partial<AngularCreateTasksParams>={}){
    params.taskSource = "angular frontend"
    super(params)
  }
}

async function getTasks(): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    // @ts-ignore
    result = [
      new AngularCreateTasksParams({taskName:"run"}),
      new AngularCreateTasksParams({taskName:"install app deps"}),
      new AngularCreateTasksParams({taskName:"check for angular updates"}),
      new AngularCreateTasksParams({taskName:"update angular"}),
      new AngularCreateTasksParams({taskName:"run compodoc"}),
      new AngularCreateTasksParams({taskName:"analyze"}),
      new AngularCreateTasksParams({taskName:"run translate"})
    ]
    .map((task)=>{
      return createTask(task)
    })
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    letDeveloperKnowAboutAnIssue(err,'Issue while loading windmillcode Angular tasks.')
  }

	return result;
}
