import * as path from "path";
import * as fs from 'fs';
import { letDeveloperKnowAboutAnIssue } from './functions';
import * as vscode from 'vscode';

async function readJsonFile(filePath:string) {
  try {
    // Use 'await' to asynchronously read the file
    const data = await fs.promises.readFile(filePath, 'utf8');

    // Parse the JSON data into a JavaScript object
    const jsonData = JSON.parse(data);

    // Now jsonData contains the contents of the JSON file
    return jsonData;
  } catch (error) {
    console.error('Error reading or parsing the file:', error);
    throw error;
  }
}

export let modifyTasksJson = async (extensionRoot:string,executable:string)=>{
  let myExtension = vscode.extensions.getExtension('windmillcode-publisher-0.windmillcode-extension-pack-0');
  let installLocation = path.normalize(extensionRoot+"/task_files")
  let tasksJsonFile = path.normalize(`${installLocation}/tasks.json`)
  tasksJsonFile = await readJsonFile(tasksJsonFile);
  letDeveloperKnowAboutAnIssue(null,tasksJsonFile)
  letDeveloperKnowAboutAnIssue(null,executable)
}
