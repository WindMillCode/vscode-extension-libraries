function Show-Menu {
    param (
        [Parameter(Mandatory = $true)]
        [string]$Prompt,

        [Parameter(Mandatory = $true)]
        [string[]]$Options,

        [Parameter()]
        [switch]$EnableOtherOption = $false
    )

    $maxOptionLength = ($Options | Measure-Object -Property Length -Maximum).Maximum
    $actualOptionsCount = $Options.Count
    if ($EnableOtherOption) {
        $actualOptionsCount++
    }


    $selectedIndex = 0
    while ($true) {
        Write-Host $Prompt
        for ($i = 0; $i -lt $actualOptionsCount; $i++) {
            if ($i -eq $Options.Count - 1 -and $EnableOtherOption) {
                $option = "OTHER"
            }
            else {
                $option = $Options[$i]
            }
            $paddedOption = $option.PadRight($maxOptionLength)

            if ($i -eq $selectedIndex) {
                Write-Host " > $($i + 1). $paddedOption"
            }
            else {
                Write-Host "   $($i + 1). $paddedOption"
            }
        }

        $keyInfo = [System.Console]::ReadKey($true)
        $key = $keyInfo.Key

        if ($key -eq 'UpArrow') {
            $selectedIndex = ($selectedIndex - 1 + $actualOptionsCount) % $actualOptionsCount
        }
        elseif ($key -eq 'DownArrow') {
            $selectedIndex = ($selectedIndex + 1) % $actualOptionsCount
        }
        elseif ($key -eq 'Enter') {
            if ($selectedIndex -eq $Options.Count - 1 -and $EnableOtherOption) {
                Write-Host
                $otherOption = Read-Host "Provide a value for OTHER:"
                return $otherOption
            }
            break
        }
        [System.Console]::Clear()
    }

    return $Options[$selectedIndex]
}


function Take-Variable-Args {


    $InnerScriptArguments = @()

    while ($true) {
        $argument = Read-Host "Enter the arguments to pass to the script (press ENTER to enter another argument, leave blank and press ENTER once done )"
        if ([string]::IsNullOrWhiteSpace($argument)) {
            break
        }
        $InnerScriptArguments += $argument
    }

    return $InnerScriptArguments
    # Invoke the inner script with the provided arguments
    # & $InnerScriptPath $InnerScriptArguments
}

function GetParamValue {
    param (
        [Parameter(Mandatory=$true)]
        [System.Reflection.ParameterInfo]$Parameter
    )

    $parameterValue = $Parameter.DefaultValue
    if (-not([System.Management.Automation.Language.NullString]::IsNullOrEmpty($parameterValue))) {
        return $parameterValue
    }
    else {
        Write-Host "Parameter value not found."
    }
}


$path = $MyInvocation.MyCommand.Path
if (!$path) {$path = $psISE.CurrentFile.Fullpath}
if ($path)  {$path = Split-Path $path -Parent}


$currentBranch = git rev-parse --abbrev-ref HEAD
