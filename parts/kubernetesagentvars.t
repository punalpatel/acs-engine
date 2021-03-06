    "{{.Name}}StorageAccountOffset": "[mul(variables('maxStorageAccountsPerAgent'),variables('{{.Name}}Index'))]",
    "{{.Name}}Count": "[parameters('{{.Name}}Count')]",
    "{{.Name}}AvailabilitySet": "[concat('{{.Name}}-availabilitySet-', variables('nameSuffix'))]",
    "{{.Name}}StorageAccountsCount": "[add(div(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')), mod(add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),2), add(mod(variables('{{.Name}}Count'), variables('maxVMsPerStorageAccount')),1)))]",
    "{{.Name}}VMNamePrefix": "[concat(variables('orchestratorName'), '-{{.Name}}-', variables('nameSuffix'), '-')]", 
    "{{.Name}}VMSize": "[parameters('{{.Name}}VMSize')]",
{{if .IsCustomVNET}}
    "{{.Name}}VnetSubnetID": "[parameters('{{.Name}}VnetSubnetID')]",
    "{{.Name}}SubnetName": "[parameters('{{.Name}}VnetSubnetID')]",
    "{{.Name}}VnetParts": "[split(parameters('{{.Name}}VnetSubnetID'),'/subnets/')]",
{{else}}
    "{{.Name}}VnetSubnetID": "[variables('vnetSubnetID')]",
    "{{.Name}}SubnetName": "[variables('subnetName')]",
{{end}}