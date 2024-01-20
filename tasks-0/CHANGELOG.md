# ChangeLog

All notable changes to the "Windmillcode Tasks Zero" extension pack will be documented in this file.

* Version updates will be based on vscode relesases
on every vscode update a new version will be release

* the software version extends the vscode patch version by 3 zeros giving us
1000 possible updates before there is an update to vscode and extends back to zero

* you would have to check the CHANGELOG for any breaking, (major), minor or patched updates which will be denoted respectively



## [1.85.1000] - 12-27-2023
* Extension made available to the public ready for use

## [1.85.1001] - 12-27-2023
* [FIX] fixed a bug with flask backend create endpoint


## [1.85.1002] - 12-27-2023
* [UPDATE] added a feature where you can view coverage info at localhost:8003 for angular_frontend_test
* [UPDATE] added a feature where you can view coverage info at localhost:8004 for flask_backend_test


## [1.85.1003] - 12-27-2023
* [PATCH] - fix bug in flask_backned_run flask_backend_test and docker_init_container trying
based on an underlying command fn from windmillcode/go_scripts package under investigation
* [UPDATE]  seperated coverage http-server to its on task flask_backend_view_coverage_info from flask_backend_test


## [1.85.1004] - 1-2-2023
* [UPDATE] - configured angular frontend run and flask backend run so a developer wont have to toggle between developer and docker development in the settings

## [1.85.1005] - 1-5-2023
* [UPDATE] - made an an update to tasks update workspace by providing additional prompts that will address devices with less capable peformance components

## [1.85.1006] - 1-5-2023
* update tasks to ensure for linux that the bash shell is used and ran in interactive mode in an attempt to source the .bashrc with important features such as the $PATH

## [1.85.1007] - 1-5-2023
* [PATCH] patch an issue with an underlying library

## [1.85.1008] - 1-6-2023
* [UPDATE] updated undelying go libraries so now output that returns a value also prints to the console

## [1.85.1009] - 1-6-2023
* [FIX] ensured the go programs build into go executables w/o error


## [1.85.1010] - 1-7-2023
8 [UPDATE] ensured the sql make db update entry starts with year so recent db snapshots appear lower in the file explorer than older ones

## [1.85.1011] - 1-11-2023
* [PATCH] - updated internal go code

## [1.85.1012] - 1-13-2023
* [FIX] - update flask_backend_create_manager to conform to the testing location for managers

## [1.85.1013] - 1-13-2023
* [PATCH] - updated sql_make_db_schema_update_entry so that the correct timestamp format comes out

## [1.85.1014] - 1-15-2023
* [PATCH] - corrected an issue with flask_backend_create_manager


## [1.85.1015] - 1-15-2023
* [UPDATE] added misc run proxy to run x amount of proxies on your running apps
the windmillcode-extension-pack-0.proxyURLs is a space seperated string for multiple entries
it optionally runs a diode tunnel for each proxy making the proxy public on the www learn more [here](https://support.diode.io/article/ss32engxlq-publish-your-local-webserver)

* [UPDATE] added customUserIsPresent option to
tasks_update_workspace_with_latest_tasks so you dont have to manully hit enter for the Windmillcode user

## [1.85.1016] - 1-15-2023
* [UPDATE] -made internal changes
* [UPDATE]- flask backend env does not print output to the console anymore

## [1.85.1017] - 1-15-2023
* [BREAKING CHANGE] -removed flask backend view coverage info
* [UPDATE] flask backend test supports both unit testing and coverage info
* [UPDATE] can leverage tasksToRunOnFolderOpen with the labels of the tasks you want to automatically run on folder open

