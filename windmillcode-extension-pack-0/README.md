# windmillcode-extension-pack-0 README
* snippets, tasks and extension packs for projects at windmillcode.
* (task extension support only for windows all platforms coming soon)
  * (plan to have go installed, tasks that support all platforms will have go installed)
## Changelog
v 0.2.0
  * support for the yarn utilities

v 0.3.0
  * support for the python & flask utilities
v 0.3.1
  * fixed issue with yarn utilities

v 0.4.0
  * support for testng, docker and misc utitlies

v 0.4.1
  * bug fix

v 0.4.2
  bug fixed to workspace file testng_e2e_run to testng_e2e_shared

v 0.4.3
  * ensured there was an option for the powershell options snippet

v 0.5.0
  * refactored all tasks to update the tasks.json if one is not there already

v 0.5.1
  * added support fore cloud firebase

v 0.5.3
  * fixed a bug

v 0.5.4
  * provided scss snippet font-face base64 template for base64 snippets

v 0.5.5
  * updated i18n scripts to to work, even through there is string and chunk option right now there is only string option

v 0.5.6
  * added support for font-face urls
  * updated package names as appropriate
  * provide for scss rgb color code

v 0.5.7
  * correct an error with wml-orig

v 0.5.8
  * updated flask route to flask app template specs
  * fixed and isses with scss for where you had to manually fixed the $i

v 0.5.9
  * fixed an error with the testng vscode script

v 0.6.0
  * added yarn_install_specifc_packages.ps1 so that a user can manage specified package in an application

v 0.7.0
  * fix angular_translate in use cases where it did not know how to properly deal with an object

v 1.0.0
  * remove angular service method from typescript snippetsschema as functionality is now the @windmillcode/angular-templates schematic package
  * complete refactor of extension tasks.json to use go instead of powershell for better cross os platform suppport

v 1.0.1
  * minor change

v 1.0.2
  * fixed a bug with firebase run

v 1.0.3
  * provided workspace folder location to flask scripts

v 2.0.0
  * added support for npm installations
  * added support where you can specfiy the installation path for npm and python application

v 2.0.1
  * fixed bug for sql_update_schema

v 2.0.2
  * corrected a bug where the check to remove package-lock.json is made for npm_install_app_deps

v 2.0.3
  * rmd fmt so flask_backend_go can work

v 2.0.4
  * took out long-hand in wml-original for color naming

V 2.1.0
  * fixed an issue concerning update angular

v 3.0.0
  * all cli scripts now can be complied into executables making task running much faster

v 3.0.1
  * removed spressman tasks-explorer takes up too much time

v 3.1.0
  * hammered out issues for mac and linux should be working now

v 3.2.0
  * angular and npm and python tasks support mutliple projects, where you can input multiple project paths and actions happen concurrently
  * add misc copy items to copy an item from location to all destination locatins

v 3.2.1
  * removed node-fetch fixed a bug

v 3.3.0
  * fixed an issue concerning other for show menu now working as intended

v 3.3.1
  * updated update_angular script based on npm recommdations

v 3.3.2
  * made a fix in go script utils

v 3.3.3
  * updated go

v 3.4.0
  * docker backup immages, allows the user to backup their images and containers to any specified location
  * added wait.groups go snippet as well

v 4.0.0
  * added flutter support
  * added dart snippets

v 4.0.1
* additional flutter dart snipperts/ tasks

v 4.0.2
* additional flutter dart snipperts/ tasks

v 4.0.3
  * flask create controller allows you to create a an endpoint and handler for a flask controller

v 4.0.4
  * added wml-spacing to make it easier to add spacing variables
V 4.1.0
* 2 powerful command line scripts
  * git clone_subdirs
      you choose a src and dest subdir location and the script pulls the src subdir
  * misc opptimize images
    * turns all imgs to jpg keeping quality and greatly reducing file size, svg excluded

v 5.0.0
make extension work for  macos

v 5.0.1
* npm added force option to install by force for npm and yarn


v 5.0.2
* added npm silience typescript which takes a regex and procceeds to comment out all files matching the regex

v 5.0.3
* andded angular frontend test

v 5.0.4
* corrected an error

v 5.0.5
* added angular_frontend_test

v 5.0.6
* corrected an error in update_workspace_with_latest tasks,
* added openAI base for angular run translate to use other server than openai server in
* updated npm to leverage os.Remove()

v 5.0.7
* added more capability for update scripts

v 5.0.8
* angular frontend run can specify configuration to choose from angular.json
* fixed an error in git_clone_subdirs
* updated snippets in typescript.json

v 5.0.9
* patch with update worksapce with latest tasks where the tasks.json file does not get created.

v 5.1.0
* removed the go_cli_library from the extension
* added choose a license vscode extension
* fixed flask_backend_dev and flask_backend_test

v 5.1.1
* update flask route in python.json to support unit test case version

v 5.1.2
* python install specifc package was overwriting requirements file prevented this

v 5.1.3
* python flask route snippet does not send back automatic 200

v 5.1.4
* python flask route snippet uses req_body instead of resp_body

v 5.1.5
* corrected issue with python install specific packages

v 5.1.6
* coorected go mod to work on mac os

v 5.1.7
* added python kwargs snippet to make it easier to grab values from kwargs

v 5.1.8
* added pytest helper fn snippets

v 5.1.9
* for flask_backend_test ensured the test case ran in an infinite loop

v 5.1.10
made install executable support any os
fault flutter test command does not show any output

v 5.1.11
* attempting to get extension to work with vscode dev containers

v 5.1.12
* making flask backend run & test universal

v 5.1.13
* all python scripts leveage pyenv global instead of pyenv shell

v 5.2.0
* added tasks_update_workspace_without_extension which will allow the user to rebuild dependencies w/o needed the extension to be installed

v 5.2.1
* small fix needed for tasks_update_workspace_without_extension

v 5.2.2
* added option to choose executable in angular frontend: analyze.
