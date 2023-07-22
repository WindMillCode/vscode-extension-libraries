/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import * as vscode from 'vscode';
import { WMLTasksJSONTaskProvider } from './TasksJSONProvider';

import { installGo } from './installGo';

let WMLDisposables: vscode.Disposable[] =[]
let WMLTaskProviders:any[] = [
	WMLTasksJSONTaskProvider
	// WMLAngularTaskProvider,
	// WMLFlaskTaskProvider,
	// WMLGitTaskProvider,
	// WMLYarnTaskProvider,
	// WMLPythonTaskProvider,
	// WMLMiscTaskProvider,
	// WMLTestNGTaskProvider,
	// WMLDockerTaskProvider,
	// WMLSQLTaskProvider
]
export async function activate(_context: vscode.ExtensionContext): Promise<void> {
	const workspaceRoot = (vscode.workspace.workspaceFolders && (vscode.workspace.workspaceFolders.length > 0))
		? vscode.workspace.workspaceFolders[0].uri.fsPath : undefined;

	vscode.window.showInformationMessage(_context.extensionPath);
	let goExecutable = await installGo(_context.extensionPath)

	if (!workspaceRoot) {
		return;
	}

	WMLDisposables = WMLTaskProviders
	.map((providerType,index0)=>{

		let providerInstance =  new providerType(workspaceRoot)
		providerInstance.goExecutable = goExecutable
		return vscode.tasks.registerTaskProvider(
			providerType.WindmillType,
			providerInstance
		);
	})

}

export function deactivate(): void {
	WMLDisposables
	.forEach((provider)=>{
		provider.dispose()
	})

}
