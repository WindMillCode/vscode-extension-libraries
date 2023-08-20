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
