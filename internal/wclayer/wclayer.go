package wclayer

//go:generate go run ../../mksyscall_windows.go -output zsyscall_windows.go -winio wclayer.go

//sys activateLayer(info *driverInfo, id string) (hr error) = vmcompute.ActivateLayer?
//sys copyLayer(info *driverInfo, srcId string, dstId string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) = vmcompute.CopyLayer?
//sys createLayer(info *driverInfo, id string, parent string) (hr error) = vmcompute.CreateLayer?
//sys createSandboxLayer(info *driverInfo, id string, parent string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) = vmcompute.CreateSandboxLayer?
//sys expandSandboxSize(info *driverInfo, id string, size uint64) (hr error) = vmcompute.ExpandSandboxSize?
//sys deactivateLayer(info *driverInfo, id string) (hr error) = vmcompute.DeactivateLayer?
//sys destroyLayer(info *driverInfo, id string) (hr error) = vmcompute.DestroyLayer?
//sys exportLayer(info *driverInfo, id string, path string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) = vmcompute.ExportLayer?
//sys getLayerMountPath(info *driverInfo, id string, length *uintptr, buffer *uint16) (hr error) = vmcompute.GetLayerMountPath?
//sys getBaseImages(buffer **uint16) (hr error) = vmcompute.GetBaseImages?
//sys importLayer(info *driverInfo, id string, path string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) = vmcompute.ImportLayer?
//sys layerExists(info *driverInfo, id string, exists *uint32) (hr error) = vmcompute.LayerExists?
//sys nameToGuid(name string, guid *GUID) (hr error) = vmcompute.NameToGuid?
//sys prepareLayer(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR) (hr error) = vmcompute.PrepareLayer?
//sys unprepareLayer(info *driverInfo, id string) (hr error) = vmcompute.UnprepareLayer?
//sys processBaseImage(path string) (hr error) = vmcompute.ProcessBaseImage?
//sys processUtilityImage(path string) (hr error) = vmcompute.ProcessUtilityImage?

//sys importLayerBegin(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) = vmcompute.ImportLayerBegin?
//sys importLayerNext(context uintptr, fileName string, fileInfo *winio.FileBasicInfo) (hr error) = vmcompute.ImportLayerNext?
//sys importLayerWrite(context uintptr, buffer []byte) (hr error) = vmcompute.ImportLayerWrite?
//sys importLayerEnd(context uintptr) (hr error) = vmcompute.ImportLayerEnd?

//sys exportLayerBegin(info *driverInfo, id string, descriptors []WC_LAYER_DESCRIPTOR, context *uintptr) (hr error) = vmcompute.ExportLayerBegin?
//sys exportLayerNext(context uintptr, fileName **uint16, fileInfo *winio.FileBasicInfo, fileSize *int64, deleted *uint32) (hr error) = vmcompute.ExportLayerNext?
//sys exportLayerRead(context uintptr, buffer []byte, bytesRead *uint32) (hr error) = vmcompute.ExportLayerRead?
//sys exportLayerEnd(context uintptr) (hr error) = vmcompute.ExportLayerEnd?
