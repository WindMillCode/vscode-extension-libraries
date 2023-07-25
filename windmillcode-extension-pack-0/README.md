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
