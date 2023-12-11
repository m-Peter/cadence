// Code generated by "stringer -type=PrimitiveStaticType -trimprefix=PrimitiveStaticType"; DO NOT EDIT.

package interpreter

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PrimitiveStaticTypeUnknown-0]
	_ = x[PrimitiveStaticTypeVoid-1]
	_ = x[PrimitiveStaticTypeAny-2]
	_ = x[PrimitiveStaticTypeNever-3]
	_ = x[PrimitiveStaticTypeAnyStruct-4]
	_ = x[PrimitiveStaticTypeAnyResource-5]
	_ = x[PrimitiveStaticTypeBool-6]
	_ = x[PrimitiveStaticTypeAddress-7]
	_ = x[PrimitiveStaticTypeString-8]
	_ = x[PrimitiveStaticTypeCharacter-9]
	_ = x[PrimitiveStaticTypeMetaType-10]
	_ = x[PrimitiveStaticTypeBlock-11]
	_ = x[PrimitiveStaticTypeAnyResourceAttachment-12]
	_ = x[PrimitiveStaticTypeAnyStructAttachment-13]
	_ = x[PrimitiveStaticTypeHashableStruct-14]
	_ = x[PrimitiveStaticTypeNumber-18]
	_ = x[PrimitiveStaticTypeSignedNumber-19]
	_ = x[PrimitiveStaticTypeInteger-24]
	_ = x[PrimitiveStaticTypeSignedInteger-25]
	_ = x[PrimitiveStaticTypeFixedPoint-30]
	_ = x[PrimitiveStaticTypeSignedFixedPoint-31]
	_ = x[PrimitiveStaticTypeInt-36]
	_ = x[PrimitiveStaticTypeInt8-37]
	_ = x[PrimitiveStaticTypeInt16-38]
	_ = x[PrimitiveStaticTypeInt32-39]
	_ = x[PrimitiveStaticTypeInt64-40]
	_ = x[PrimitiveStaticTypeInt128-41]
	_ = x[PrimitiveStaticTypeInt256-42]
	_ = x[PrimitiveStaticTypeUInt-44]
	_ = x[PrimitiveStaticTypeUInt8-45]
	_ = x[PrimitiveStaticTypeUInt16-46]
	_ = x[PrimitiveStaticTypeUInt32-47]
	_ = x[PrimitiveStaticTypeUInt64-48]
	_ = x[PrimitiveStaticTypeUInt128-49]
	_ = x[PrimitiveStaticTypeUInt256-50]
	_ = x[PrimitiveStaticTypeWord8-53]
	_ = x[PrimitiveStaticTypeWord16-54]
	_ = x[PrimitiveStaticTypeWord32-55]
	_ = x[PrimitiveStaticTypeWord64-56]
	_ = x[PrimitiveStaticTypeWord128-57]
	_ = x[PrimitiveStaticTypeWord256-58]
	_ = x[PrimitiveStaticTypeFix64-64]
	_ = x[PrimitiveStaticTypeUFix64-72]
	_ = x[PrimitiveStaticTypePath-76]
	_ = x[PrimitiveStaticTypeCapability-77]
	_ = x[PrimitiveStaticTypeStoragePath-78]
	_ = x[PrimitiveStaticTypeCapabilityPath-79]
	_ = x[PrimitiveStaticTypePublicPath-80]
	_ = x[PrimitiveStaticTypePrivatePath-81]
	_ = x[PrimitiveStaticTypeAuthAccount-90]
	_ = x[PrimitiveStaticTypePublicAccount-91]
	_ = x[PrimitiveStaticTypeDeployedContract-92]
	_ = x[PrimitiveStaticTypeAuthAccountContracts-93]
	_ = x[PrimitiveStaticTypePublicAccountContracts-94]
	_ = x[PrimitiveStaticTypeAuthAccountKeys-95]
	_ = x[PrimitiveStaticTypePublicAccountKeys-96]
	_ = x[PrimitiveStaticTypeAccountKey-97]
	_ = x[PrimitiveStaticTypeAuthAccountInbox-98]
	_ = x[PrimitiveStaticTypeStorageCapabilityController-99]
	_ = x[PrimitiveStaticTypeAccountCapabilityController-100]
	_ = x[PrimitiveStaticTypeAuthAccountStorageCapabilities-101]
	_ = x[PrimitiveStaticTypeAuthAccountAccountCapabilities-102]
	_ = x[PrimitiveStaticTypeAuthAccountCapabilities-103]
	_ = x[PrimitiveStaticTypePublicAccountCapabilities-104]
	_ = x[PrimitiveStaticTypeAccount-105]
	_ = x[PrimitiveStaticTypeAccount_Contracts-106]
	_ = x[PrimitiveStaticTypeAccount_Keys-107]
	_ = x[PrimitiveStaticTypeAccount_Inbox-108]
	_ = x[PrimitiveStaticTypeAccount_StorageCapabilities-109]
	_ = x[PrimitiveStaticTypeAccount_AccountCapabilities-110]
	_ = x[PrimitiveStaticTypeAccount_Capabilities-111]
	_ = x[PrimitiveStaticTypeAccount_Storage-112]
	_ = x[PrimitiveStaticTypeMutate-118]
	_ = x[PrimitiveStaticTypeInsert-119]
	_ = x[PrimitiveStaticTypeRemove-120]
	_ = x[PrimitiveStaticTypeIdentity-121]
	_ = x[PrimitiveStaticTypeStorage-125]
	_ = x[PrimitiveStaticTypeSaveValue-126]
	_ = x[PrimitiveStaticTypeLoadValue-127]
	_ = x[PrimitiveStaticTypeCopyValue-128]
	_ = x[PrimitiveStaticTypeBorrowValue-129]
	_ = x[PrimitiveStaticTypeContracts-130]
	_ = x[PrimitiveStaticTypeAddContract-131]
	_ = x[PrimitiveStaticTypeUpdateContract-132]
	_ = x[PrimitiveStaticTypeRemoveContract-133]
	_ = x[PrimitiveStaticTypeKeys-134]
	_ = x[PrimitiveStaticTypeAddKey-135]
	_ = x[PrimitiveStaticTypeRevokeKey-136]
	_ = x[PrimitiveStaticTypeInbox-137]
	_ = x[PrimitiveStaticTypePublishInboxCapability-138]
	_ = x[PrimitiveStaticTypeUnpublishInboxCapability-139]
	_ = x[PrimitiveStaticTypeClaimInboxCapability-140]
	_ = x[PrimitiveStaticTypeCapabilities-141]
	_ = x[PrimitiveStaticTypeStorageCapabilities-142]
	_ = x[PrimitiveStaticTypeAccountCapabilities-143]
	_ = x[PrimitiveStaticTypePublishCapability-144]
	_ = x[PrimitiveStaticTypeUnpublishCapability-145]
	_ = x[PrimitiveStaticTypeGetStorageCapabilityController-146]
	_ = x[PrimitiveStaticTypeIssueStorageCapabilityController-147]
	_ = x[PrimitiveStaticTypeGetAccountCapabilityController-148]
	_ = x[PrimitiveStaticTypeIssueAccountCapabilityController-149]
	_ = x[PrimitiveStaticTypeCapabilitiesMapping-150]
	_ = x[PrimitiveStaticTypeAccountMapping-151]
	_ = x[PrimitiveStaticType_Count-152]
}

const _PrimitiveStaticType_name = "UnknownVoidAnyNeverAnyStructAnyResourceBoolAddressStringCharacterMetaTypeBlockAnyResourceAttachmentAnyStructAttachmentHashableStructNumberSignedNumberIntegerSignedIntegerFixedPointSignedFixedPointIntInt8Int16Int32Int64Int128Int256UIntUInt8UInt16UInt32UInt64UInt128UInt256Word8Word16Word32Word64Word128Word256Fix64UFix64PathCapabilityStoragePathCapabilityPathPublicPathPrivatePathAuthAccountPublicAccountDeployedContractAuthAccountContractsPublicAccountContractsAuthAccountKeysPublicAccountKeysAccountKeyAuthAccountInboxStorageCapabilityControllerAccountCapabilityControllerAuthAccountStorageCapabilitiesAuthAccountAccountCapabilitiesAuthAccountCapabilitiesPublicAccountCapabilitiesAccountAccount_ContractsAccount_KeysAccount_InboxAccount_StorageCapabilitiesAccount_AccountCapabilitiesAccount_CapabilitiesAccount_StorageMutateInsertRemoveIdentityStorageSaveValueLoadValueCopyValueBorrowValueContractsAddContractUpdateContractRemoveContractKeysAddKeyRevokeKeyInboxPublishInboxCapabilityUnpublishInboxCapabilityClaimInboxCapabilityCapabilitiesStorageCapabilitiesAccountCapabilitiesPublishCapabilityUnpublishCapabilityGetStorageCapabilityControllerIssueStorageCapabilityControllerGetAccountCapabilityControllerIssueAccountCapabilityControllerCapabilitiesMappingAccountMapping_Count"

var _PrimitiveStaticType_map = map[PrimitiveStaticType]string{
	0:   _PrimitiveStaticType_name[0:7],
	1:   _PrimitiveStaticType_name[7:11],
	2:   _PrimitiveStaticType_name[11:14],
	3:   _PrimitiveStaticType_name[14:19],
	4:   _PrimitiveStaticType_name[19:28],
	5:   _PrimitiveStaticType_name[28:39],
	6:   _PrimitiveStaticType_name[39:43],
	7:   _PrimitiveStaticType_name[43:50],
	8:   _PrimitiveStaticType_name[50:56],
	9:   _PrimitiveStaticType_name[56:65],
	10:  _PrimitiveStaticType_name[65:73],
	11:  _PrimitiveStaticType_name[73:78],
	12:  _PrimitiveStaticType_name[78:99],
	13:  _PrimitiveStaticType_name[99:118],
	14:  _PrimitiveStaticType_name[118:132],
	18:  _PrimitiveStaticType_name[132:138],
	19:  _PrimitiveStaticType_name[138:150],
	24:  _PrimitiveStaticType_name[150:157],
	25:  _PrimitiveStaticType_name[157:170],
	30:  _PrimitiveStaticType_name[170:180],
	31:  _PrimitiveStaticType_name[180:196],
	36:  _PrimitiveStaticType_name[196:199],
	37:  _PrimitiveStaticType_name[199:203],
	38:  _PrimitiveStaticType_name[203:208],
	39:  _PrimitiveStaticType_name[208:213],
	40:  _PrimitiveStaticType_name[213:218],
	41:  _PrimitiveStaticType_name[218:224],
	42:  _PrimitiveStaticType_name[224:230],
	44:  _PrimitiveStaticType_name[230:234],
	45:  _PrimitiveStaticType_name[234:239],
	46:  _PrimitiveStaticType_name[239:245],
	47:  _PrimitiveStaticType_name[245:251],
	48:  _PrimitiveStaticType_name[251:257],
	49:  _PrimitiveStaticType_name[257:264],
	50:  _PrimitiveStaticType_name[264:271],
	53:  _PrimitiveStaticType_name[271:276],
	54:  _PrimitiveStaticType_name[276:282],
	55:  _PrimitiveStaticType_name[282:288],
	56:  _PrimitiveStaticType_name[288:294],
	57:  _PrimitiveStaticType_name[294:301],
	58:  _PrimitiveStaticType_name[301:308],
	64:  _PrimitiveStaticType_name[308:313],
	72:  _PrimitiveStaticType_name[313:319],
	76:  _PrimitiveStaticType_name[319:323],
	77:  _PrimitiveStaticType_name[323:333],
	78:  _PrimitiveStaticType_name[333:344],
	79:  _PrimitiveStaticType_name[344:358],
	80:  _PrimitiveStaticType_name[358:368],
	81:  _PrimitiveStaticType_name[368:379],
	90:  _PrimitiveStaticType_name[379:390],
	91:  _PrimitiveStaticType_name[390:403],
	92:  _PrimitiveStaticType_name[403:419],
	93:  _PrimitiveStaticType_name[419:439],
	94:  _PrimitiveStaticType_name[439:461],
	95:  _PrimitiveStaticType_name[461:476],
	96:  _PrimitiveStaticType_name[476:493],
	97:  _PrimitiveStaticType_name[493:503],
	98:  _PrimitiveStaticType_name[503:519],
	99:  _PrimitiveStaticType_name[519:546],
	100: _PrimitiveStaticType_name[546:573],
	101: _PrimitiveStaticType_name[573:603],
	102: _PrimitiveStaticType_name[603:633],
	103: _PrimitiveStaticType_name[633:656],
	104: _PrimitiveStaticType_name[656:681],
	105: _PrimitiveStaticType_name[681:688],
	106: _PrimitiveStaticType_name[688:705],
	107: _PrimitiveStaticType_name[705:717],
	108: _PrimitiveStaticType_name[717:730],
	109: _PrimitiveStaticType_name[730:757],
	110: _PrimitiveStaticType_name[757:784],
	111: _PrimitiveStaticType_name[784:804],
	112: _PrimitiveStaticType_name[804:819],
	118: _PrimitiveStaticType_name[819:825],
	119: _PrimitiveStaticType_name[825:831],
	120: _PrimitiveStaticType_name[831:837],
	121: _PrimitiveStaticType_name[837:845],
	125: _PrimitiveStaticType_name[845:852],
	126: _PrimitiveStaticType_name[852:861],
	127: _PrimitiveStaticType_name[861:870],
	128: _PrimitiveStaticType_name[870:879],
	129: _PrimitiveStaticType_name[879:890],
	130: _PrimitiveStaticType_name[890:899],
	131: _PrimitiveStaticType_name[899:910],
	132: _PrimitiveStaticType_name[910:924],
	133: _PrimitiveStaticType_name[924:938],
	134: _PrimitiveStaticType_name[938:942],
	135: _PrimitiveStaticType_name[942:948],
	136: _PrimitiveStaticType_name[948:957],
	137: _PrimitiveStaticType_name[957:962],
	138: _PrimitiveStaticType_name[962:984],
	139: _PrimitiveStaticType_name[984:1008],
	140: _PrimitiveStaticType_name[1008:1028],
	141: _PrimitiveStaticType_name[1028:1040],
	142: _PrimitiveStaticType_name[1040:1059],
	143: _PrimitiveStaticType_name[1059:1078],
	144: _PrimitiveStaticType_name[1078:1095],
	145: _PrimitiveStaticType_name[1095:1114],
	146: _PrimitiveStaticType_name[1114:1144],
	147: _PrimitiveStaticType_name[1144:1176],
	148: _PrimitiveStaticType_name[1176:1206],
	149: _PrimitiveStaticType_name[1206:1238],
	150: _PrimitiveStaticType_name[1238:1257],
	151: _PrimitiveStaticType_name[1257:1271],
	152: _PrimitiveStaticType_name[1271:1277],
}

func (i PrimitiveStaticType) String() string {
	if str, ok := _PrimitiveStaticType_map[i]; ok {
		return str
	}
	return "PrimitiveStaticType(" + strconv.FormatInt(int64(i), 10) + ")"
}
