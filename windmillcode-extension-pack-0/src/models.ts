import * as vscode from 'vscode';

export interface WMLTaskDefinition extends vscode.TaskDefinition {

	task: string;
}
export enum OperatingSystem  {
  AIX = 'AIX',
  MACOS = 'darwin',
  FREEBSD = 'freebsd',
  LINUX = 'linux',
  OPENBSD = 'openbsd',
  SUNOS = 'sunos',
  WINDOWS = 'win32',
};


