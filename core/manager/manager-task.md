1. Source pulling, reporting, data retention
2. Pre- and Post-scan actions, branching projects, scheduled tasks/scans

The Manager is responsible for picking up and completing the source pulling, and doing some preprocessing of the source code including counting LOC, enforcing exclusions, calculating hashes to understand if code has changed, beautifying minified, and copying source to Scan_dir (if there are changes).


The Manager handles all incoming scan requests and is responsible for requesting source pulling and creating scan requests in the database

Manger log [core/logs/manager/etc.../
