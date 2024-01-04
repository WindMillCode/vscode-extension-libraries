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
