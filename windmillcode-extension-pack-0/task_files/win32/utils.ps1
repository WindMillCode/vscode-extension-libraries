function Show-Menu {
    param (
        [Parameter(Mandatory = $true)]
        [string]$Prompt,

        [Parameter(Mandatory = $true)]
        [string[]]$Options
    )

    # Determine the maximum option length for padding
    $maxOptionLength = ($Options | Measure-Object -Property Length -Maximum).Maximum

    # Display the menu options
    Write-Host $Prompt
    for ($i = 0; $i -lt $Options.Count; $i++) {
        $option = $Options[$i]
        $paddedOption = $option.PadRight($maxOptionLength)

        # Check if the current option is selected
        if ($i -eq 0) {
            Write-Host " > $($i + 1). $paddedOption"
        }
        else {
            Write-Host "   $($i + 1). $paddedOption"
        }
    }

    # Prompt the user for selection
    $selectedIndex = 0
    while ($true) {
        $keyInfo = [System.Console]::ReadKey($true)
        $key = $keyInfo.Key

        if ($key -eq 'UpArrow') {
            $selectedIndex = ($selectedIndex - 1) % $Options.Count
        }
        elseif ($key -eq 'DownArrow') {
            $selectedIndex = ($selectedIndex + 1) % $Options.Count
        }
        elseif ($key -eq 'Enter') {
            break
        }

        # Clear the console and redraw menu with updated selection
        [System.Console]::Clear()
        Write-Host $Prompt
        for ($i = 0; $i -lt $Options.Count; $i++) {
            $option = $Options[$i]
            $paddedOption = $option.PadRight($maxOptionLength)

            # Check if the current option is selected
            if ($i -eq $selectedIndex) {
                Write-Host " > $($i + 1). $paddedOption"
            }
            else {
                Write-Host "   $($i + 1). $paddedOption"
            }
        }
    }

    # Return the selected option
    return $Options[$selectedIndex]
}



$path = $MyInvocation.MyCommand.Path
if (!$path) {$path = $psISE.CurrentFile.Fullpath}
if ($path)  {$path = Split-Path $path -Parent}
