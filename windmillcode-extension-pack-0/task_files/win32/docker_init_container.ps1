Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $initScript ="",
    [string[]] $initScriptArgs=@()
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

if ( $initScript -eq "") {
    $initScript  = Read-Host -Prompt "docker init script to run relative to workspace root (leave empty for the default )"
    if($initScript -eq ""){
        $initScript ="ignore\Local\docker_init_container.ps1"
    }
}

if( $initScriptArgs.Count -eq 0){
    $initScriptArgs = Take-Variable-Args
}

if($initScriptArgs -eq ""){
    & $initScript $initScriptArgs
}
else{
    & $initScript
}
