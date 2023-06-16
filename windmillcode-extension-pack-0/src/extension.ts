/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/
import * as vscode from 'vscode';
import { WMLGitTaskProvider } from './WMLGitTaskProvider';
import { getOutputChannel } from './functions';

let WMLDisposables: vscode.Disposable[] =[]
let WMLTaskProviders:any[] = [
	WMLGitTaskProvider
]
export function activate(_context: vscode.ExtensionContext): void {
	const workspaceRoot = (vscode.workspace.workspaceFolders && (vscode.workspace.workspaceFolders.length > 0))
		? vscode.workspace.workspaceFolders[0].uri.fsPath : undefined;
	if (!workspaceRoot) {
		return;
	}

	WMLDisposables = WMLTaskProviders
	.map((providerType,index0)=>{

		return vscode.tasks.registerTaskProvider(
			providerType.WindmillType,
			 new providerType(workspaceRoot)
		);
	})

}

export function deactivate(): void {
	WMLDisposables
	.forEach((provider)=>{
		provider.dispose()
	})

}
