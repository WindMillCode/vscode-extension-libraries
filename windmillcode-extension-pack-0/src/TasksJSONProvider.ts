/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as vscode from 'vscode';
import { CreateTaskParams, createTask, notifyDeveloper } from './functions';
import * as path from "path"



export class WMLTasksJSONTaskProvider implements vscode.TaskProvider {
	static WindmillType = 'windmillcode';
  goExecutable!:string
	constructor(workspaceRoot: string) {
	}

	public provideTasks(): Thenable<vscode.Task[]> | undefined {

		return getTasks(this.goExecutable);
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

async function getTasks(goExecutable:string): Promise<vscode.Task[]> {

	let result: vscode.Task[]  = [];
  try {

    notifyDeveloper(null,goExecutable)
    // @ts-ignore
    result = [
      new TasksJSONCreateTasksParams({
        executable:goExecutable,
        taskName:"update workspace with latest tasks"
      }),
    ]
    .map((taskParams,index0)=>{
      let newTask = createTask(taskParams) as any
      if(index0 === 0){
        newTask.execution.commandLine += `
        ${taskParams.extensionFolder}
        ${path.join("task_files","tasks.json")}
        ${goExecutable}
        `.split("\n").join(" ")
      }
      return newTask
    })
    .filter((task)=>{
      return task instanceof vscode.Task
    })


  } catch (err: any) {
    notifyDeveloper(err,'Issue while loading windmillcode TasksJSON tasks.')
  }

	return result;
}
