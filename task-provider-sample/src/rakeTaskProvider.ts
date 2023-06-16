/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import * as path from 'path';
import * as fs from 'fs';
import * as cp from 'child_process';
import * as vscode from 'vscode';
import { error } from 'console';

let _channel: vscode.OutputChannel;
function getOutputChannel(): vscode.OutputChannel {
	if (!_channel) {
		_channel = vscode.window.createOutputChannel('Rake Auto Detection');
	}
	return _channel;
}


export class RakeTaskProvider implements vscode.TaskProvider {
	static RakeType = 'rake';
	private rakePromise: Thenable<vscode.Task[]> | undefined = undefined;

	constructor(workspaceRoot: string) {
		const pattern = path.join(workspaceRoot, 'Rakefile');
		const fileWatcher = vscode.workspace.createFileSystemWatcher(pattern);
		fileWatcher.onDidChange(() => this.rakePromise = undefined);
		fileWatcher.onDidCreate(() => this.rakePromise = undefined);
		fileWatcher.onDidDelete(() => this.rakePromise = undefined);
	}

	public provideTasks(): Thenable<vscode.Task[]> | undefined {
		if (!this.rakePromise) {
			this.rakePromise = getRakeTasks();
		}
		return this.rakePromise;
	}

	// public resolveTask(_task: vscode.Task): vscode.Task | undefined {
	// 	const task = _task.definition.task;
	// 	let channel = getOutputChannel()
	// 	channel.appendLine(String(1))
	// 	// A Rake task consists of a task and an optional file as specified in RakeTaskDefinition
	// 	// Make sure that this looks like a Rake task by checking that there is a task.
	// 	if (task) {
	// 		// resolveTask requires that the same definition object be used.
	// 		const definition: RakeTaskDefinition = <any>_task.definition;
	// 		return new vscode.Task(
	// 			definition,
	// 			 _task.scope ?? vscode.TaskScope.Workspace,
	// 			  definition.task,
	// 				 'rake',
	// 				  new vscode.ShellExecution(`rake ${definition.task}`));
	// 	}
	// 	return undefined;
	// }

	public resolveTask(_task: vscode.Task): vscode.Task | undefined {
		return new vscode.Task(
			_task.definition,
			vscode.TaskScope.Workspace,
			"see me?",
			'windmillcode',
			new vscode.ShellExecution(`ls`));
	}
}



function exec(command: string, options: cp.ExecOptions): Promise<{ stdout: string; stderr: string }> {
	return new Promise<{ stdout: string; stderr: string }>((resolve, reject) => {
		cp.exec(command, options, (error, stdout, stderr) => {
			if (error) {
				reject({ error, stdout, stderr });
			}
			resolve({ stdout, stderr });
		});
	});
}


interface RakeTaskDefinition extends vscode.TaskDefinition {
	/**
	 * The task name
	 */
	task: string;

	/**
	 * The rake file containing the task
	 */
	file?: string;
}

const buildNames: string[] = ['build', 'compile', 'watch'];
function isBuildTask(name: string): boolean {
	return true
}

const testNames: string[] = ['test'];
function isTestTask(name: string): boolean {
	return true
}

async function getRakeTasks(): Promise<vscode.Task[]> {
	const workspaceFolders = vscode.workspace.workspaceFolders;
	const result: vscode.Task[] = [];
	const channel = getOutputChannel();
	channel.appendLine(result.length.toString())
	if (!workspaceFolders || workspaceFolders.length === 0) {
		return result;
	}
	for (const workspaceFolder of workspaceFolders) {
		const folderString = workspaceFolder.uri.fsPath;


		const commandLine = 'rake';
		try {
			const { stdout, stderr } = await exec(commandLine, { cwd: folderString });
			if (stderr && stderr.length > 0) {
				getOutputChannel().appendLine(stderr);
				getOutputChannel().show(true);
			}
			if (stdout) {
				const lines = stdout.split(/\r{0,1}\n/);
				for (const line of lines) {
					if (line.length === 0) {
						continue;
					}
					const regExp = /rake\s(.*)#/;
					const matches = regExp.exec(line);
					if (matches && matches.length === 2) {
						const taskName = matches[1].trim();
						const kind: RakeTaskDefinition = {
							type: 'rake',
							task: taskName
						};
						const task = new vscode.Task(kind, workspaceFolder, taskName, 'rake', new vscode.ShellExecution(`rake ${taskName}`));
						result.push(task);
						const lowerCaseLine = line.toLowerCase();
						if (isBuildTask(lowerCaseLine)) {
							task.group = vscode.TaskGroup.Build;
						} else if (isTestTask(lowerCaseLine)) {
							task.group = vscode.TaskGroup.Test;
						}
					}
				}
			}
		} catch (err: any) {
			const channel = getOutputChannel();
			if (err.stderr) {
				channel.appendLine(err.stderr);
			}
			if (err.stdout) {
				channel.appendLine(err.stdout);
			}
			channel.appendLine('Could not auto detect rake tasks.');
			channel.show(true);
		}
	}
	return result;
}
