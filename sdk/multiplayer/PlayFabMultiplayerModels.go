package multiplayer

import "time"
                    
// AssetReference 
type AssetReferenceModel struct {
    // FileName the asset's file name. This is a filename with the .zip, .tar, or .tar.gz extension.
    FileName string `json:"FileName,omitempty"`
    // MountPath the asset's mount path.
    MountPath string `json:"MountPath,omitempty"`
}

// AssetReferenceParams 
type AssetReferenceParamsModel struct {
    // FileName the asset's file name.
    FileName string `json:"FileName,omitempty"`
    // MountPath the asset's mount path.
    MountPath string `json:"MountPath,omitempty"`
}

// AssetSummary 
type AssetSummaryModel struct {
    // FileName the asset's file name. This is a filename with the .zip, .tar, or .tar.gz extension.
    FileName string `json:"FileName,omitempty"`
    // Metadata the metadata associated with the asset.
    Metadata map[string]string `json:"Metadata,omitempty"`
}

// AttributeMergeFunction 
type AttributeMergeFunction string
  
const (
     AttributeMergeFunctionMin AttributeMergeFunction = "Min"
     AttributeMergeFunctionMax AttributeMergeFunction = "Max"
     AttributeMergeFunctionAverage AttributeMergeFunction = "Average"
      )
// AttributeNotSpecifiedBehavior 
type AttributeNotSpecifiedBehavior string
  
const (
     AttributeNotSpecifiedBehaviorUseDefault AttributeNotSpecifiedBehavior = "UseDefault"
     AttributeNotSpecifiedBehaviorMatchAny AttributeNotSpecifiedBehavior = "MatchAny"
      )
// AttributeSource 
type AttributeSource string
  
const (
     AttributeSourceUser AttributeSource = "User"
     AttributeSourcePlayerEntity AttributeSource = "PlayerEntity"
      )
// AzureRegion 
type AzureRegion string
  
const (
     AzureRegionAustraliaEast AzureRegion = "AustraliaEast"
     AzureRegionAustraliaSoutheast AzureRegion = "AustraliaSoutheast"
     AzureRegionBrazilSouth AzureRegion = "BrazilSouth"
     AzureRegionCentralUs AzureRegion = "CentralUs"
     AzureRegionEastAsia AzureRegion = "EastAsia"
     AzureRegionEastUs AzureRegion = "EastUs"
     AzureRegionEastUs2 AzureRegion = "EastUs2"
     AzureRegionJapanEast AzureRegion = "JapanEast"
     AzureRegionJapanWest AzureRegion = "JapanWest"
     AzureRegionNorthCentralUs AzureRegion = "NorthCentralUs"
     AzureRegionNorthEurope AzureRegion = "NorthEurope"
     AzureRegionSouthCentralUs AzureRegion = "SouthCentralUs"
     AzureRegionSoutheastAsia AzureRegion = "SoutheastAsia"
     AzureRegionWestEurope AzureRegion = "WestEurope"
     AzureRegionWestUs AzureRegion = "WestUs"
     AzureRegionChinaEast2 AzureRegion = "ChinaEast2"
     AzureRegionChinaNorth2 AzureRegion = "ChinaNorth2"
     AzureRegionSouthAfricaNorth AzureRegion = "SouthAfricaNorth"
      )
// AzureVmFamily 
type AzureVmFamily string
  
const (
     AzureVmFamilyA AzureVmFamily = "A"
     AzureVmFamilyAv2 AzureVmFamily = "Av2"
     AzureVmFamilyDv2 AzureVmFamily = "Dv2"
     AzureVmFamilyF AzureVmFamily = "F"
     AzureVmFamilyFsv2 AzureVmFamily = "Fsv2"
      )
// AzureVmSize 
type AzureVmSize string
  
const (
     AzureVmSizeStandard_D1_v2 AzureVmSize = "Standard_D1_v2"
     AzureVmSizeStandard_D2_v2 AzureVmSize = "Standard_D2_v2"
     AzureVmSizeStandard_D3_v2 AzureVmSize = "Standard_D3_v2"
     AzureVmSizeStandard_D4_v2 AzureVmSize = "Standard_D4_v2"
     AzureVmSizeStandard_D5_v2 AzureVmSize = "Standard_D5_v2"
     AzureVmSizeStandard_A1_v2 AzureVmSize = "Standard_A1_v2"
     AzureVmSizeStandard_A2_v2 AzureVmSize = "Standard_A2_v2"
     AzureVmSizeStandard_A4_v2 AzureVmSize = "Standard_A4_v2"
     AzureVmSizeStandard_A8_v2 AzureVmSize = "Standard_A8_v2"
     AzureVmSizeStandard_F1 AzureVmSize = "Standard_F1"
     AzureVmSizeStandard_F2 AzureVmSize = "Standard_F2"
     AzureVmSizeStandard_F4 AzureVmSize = "Standard_F4"
     AzureVmSizeStandard_F8 AzureVmSize = "Standard_F8"
     AzureVmSizeStandard_F16 AzureVmSize = "Standard_F16"
     AzureVmSizeStandard_F2s_v2 AzureVmSize = "Standard_F2s_v2"
     AzureVmSizeStandard_F4s_v2 AzureVmSize = "Standard_F4s_v2"
     AzureVmSizeStandard_F8s_v2 AzureVmSize = "Standard_F8s_v2"
     AzureVmSizeStandard_F16s_v2 AzureVmSize = "Standard_F16s_v2"
     AzureVmSizeStandard_A1 AzureVmSize = "Standard_A1"
     AzureVmSizeStandard_A2 AzureVmSize = "Standard_A2"
     AzureVmSizeStandard_A3 AzureVmSize = "Standard_A3"
     AzureVmSizeStandard_A4 AzureVmSize = "Standard_A4"
      )
// BuildRegion 
type BuildRegionModel struct {
    // CurrentServerStats the current multiplayer server stats for the region.
    CurrentServerStats *CurrentServerStatsModel `json:"CurrentServerStats,omitempty"`
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // Region the build region.
    Region AzureRegion `json:"Region,omitempty"`
    // StandbyServers the number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
    // Status the status of multiplayer servers in the build region. Valid values are - Unknown, Initialized, Deploying, Deployed,
// Unhealthy.
    Status string `json:"Status,omitempty"`
}

// BuildRegionParams 
type BuildRegionParamsModel struct {
    // MaxServers the maximum number of multiplayer servers for the region.
    MaxServers int32 `json:"MaxServers,omitempty"`
    // Region the build region.
    Region AzureRegion `json:"Region,omitempty"`
    // StandbyServers the number of standby multiplayer servers for the region.
    StandbyServers int32 `json:"StandbyServers,omitempty"`
}

// BuildSummary 
type BuildSummaryModel struct {
    // BuildId the guid string build ID of the build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // RegionConfigurations the configuration and status for each region in the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
}

// CancelAllMatchmakingTicketsForPlayerRequest cancels all tickets of which the player is a member in a given queue that are not cancelled or matched. This API is
// useful if you lose track of what tickets the player is a member of (if the title crashes for instance) and want to
// "reset". The Entity field is optional if the caller is a player and defaults to that player. Players may not cancel
// tickets for other people. The Entity field is required if the caller is a server (authenticated as the title).
type CancelAllMatchmakingTicketsForPlayerRequestModel struct {
    // Entity the entity key of the player whose tickets should be canceled.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue from which a player's tickets should be canceled.
    QueueName string `json:"QueueName,omitempty"`
}

// CancelAllMatchmakingTicketsForPlayerResult 
type CancelAllMatchmakingTicketsForPlayerResultModel struct {
}

// CancelAllServerBackfillTicketsForPlayerRequest cancels all backfill tickets of which the player is a member in a given queue that are not cancelled or matched. This
// API is useful if you lose track of what tickets the player is a member of (if the server crashes for instance) and want
// to "reset".
type CancelAllServerBackfillTicketsForPlayerRequestModel struct {
    // Entity the entity key of the player whose backfill tickets should be canceled.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue from which a player's backfill tickets should be canceled.
    QueueName string `json:"QueueName,omitempty"`
}

// CancelAllServerBackfillTicketsForPlayerResult 
type CancelAllServerBackfillTicketsForPlayerResultModel struct {
}

// CancellationReason 
type CancellationReason string
  
const (
     CancellationReasonRequested CancellationReason = "Requested"
     CancellationReasonInternal CancellationReason = "Internal"
     CancellationReasonTimeout CancellationReason = "Timeout"
      )
// CancelMatchmakingTicketRequest only servers and ticket members can cancel a ticket. The ticket can be in five different states when it is cancelled. 1:
// the ticket is waiting for members to join it, and it has not started matching. If the ticket is cancelled at this stage,
// it will never match. 2: the ticket is matching. If the ticket is cancelled, it will stop matching. 3: the ticket is
// matched. A matched ticket cannot be cancelled. 4: the ticket is already cancelled and nothing happens. 5: the ticket is
// waiting for a server. If the ticket is cancelled, server allocation will be stopped. A server may still be allocated due
// to a race condition, but that will not be reflected in the ticket. There may be race conditions between the ticket
// getting matched and the client making a cancellation request. The client must handle the possibility that the cancel
// request fails if a match is found before the cancellation request is processed. We do not allow resubmitting a cancelled
// ticket because players must consent to enter matchmaking again. Create a new ticket instead.
type CancelMatchmakingTicketRequestModel struct {
    // QueueName the name of the queue the ticket is in.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CancelMatchmakingTicketResult 
type CancelMatchmakingTicketResultModel struct {
}

// CancelServerBackfillTicketRequest only servers can cancel a backfill ticket. The ticket can be in three different states when it is cancelled. 1: the
// ticket is matching. If the ticket is cancelled, it will stop matching. 2: the ticket is matched. A matched ticket cannot
// be cancelled. 3: the ticket is already cancelled and nothing happens. There may be race conditions between the ticket
// getting matched and the server making a cancellation request. The server must handle the possibility that the cancel
// request fails if a match is found before the cancellation request is processed. We do not allow resubmitting a cancelled
// ticket. Create a new ticket instead.
type CancelServerBackfillTicketRequestModel struct {
    // QueueName the name of the queue the ticket is in.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CancelServerBackfillTicketResult 
type CancelServerBackfillTicketResultModel struct {
}

// Certificate 
type CertificateModel struct {
    // Base64EncodedValue base64 encoded string contents of the certificate.
    Base64EncodedValue string `json:"Base64EncodedValue,omitempty"`
    // Name a name for the certificate. This is used to reference certificates in build configurations.
    Name string `json:"Name,omitempty"`
    // Password if required for your PFX certificate, use this field to provide a password that will be used to install the certificate
// on the container.
    Password string `json:"Password,omitempty"`
}

// CertificateSummary 
type CertificateSummaryModel struct {
    // Name the name of the certificate.
    Name string `json:"Name,omitempty"`
    // Thumbprint the thumbprint for the certificate.
    Thumbprint string `json:"Thumbprint,omitempty"`
}

// CognitiveServicesType 
type CognitiveServicesType string
  
const (
     CognitiveServicesTypeSpeechToText CognitiveServicesType = "SpeechToText"
     CognitiveServicesTypeTextToSpeech CognitiveServicesType = "TextToSpeech"
      )
// ConnectedPlayer 
type ConnectedPlayerModel struct {
    // PlayerId the player ID of the player connected to the multiplayer server.
    PlayerId string `json:"PlayerId,omitempty"`
}

// ContainerFlavor 
type ContainerFlavor string
  
const (
     ContainerFlavorManagedWindowsServerCore ContainerFlavor = "ManagedWindowsServerCore"
     ContainerFlavorCustomLinux ContainerFlavor = "CustomLinux"
     ContainerFlavorManagedWindowsServerCorePreview ContainerFlavor = "ManagedWindowsServerCorePreview"
      )
// ContainerImageReference 
type ContainerImageReferenceModel struct {
    // ImageName the container image name.
    ImageName string `json:"ImageName,omitempty"`
    // Tag the container tag.
    Tag string `json:"Tag,omitempty"`
}

// CoreCapacity 
type CoreCapacityModel struct {
    // Available the available core capacity for the (Region, VmFamily)
    Available int32 `json:"Available,omitempty"`
    // Region the AzureRegion
    Region AzureRegion `json:"Region,omitempty"`
    // Total the total core capacity for the (Region, VmFamily)
    Total int32 `json:"Total,omitempty"`
    // VmFamily the AzureVmFamily
    VmFamily AzureVmFamily `json:"VmFamily,omitempty"`
}

// CreateBuildWithCustomContainerRequest creates a multiplayer server build with a custom container and returns information about the build creation request.
type CreateBuildWithCustomContainerRequestModel struct {
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerImageReference the container reference, consisting of the image name and tag.
    ContainerImageReference *ContainerImageReferenceModel `json:"ContainerImageReference,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithCustomContainerResponse 
type CreateBuildWithCustomContainerResponseModel struct {
    // BuildId the guid string build ID. Must be unique for every build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container of the build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // CustomGameContainerImage the custom game container image reference information.
    CustomGameContainerImage *ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithManagedContainerRequest creates a multiplayer server build with a managed container and returns information about the build creation request.
type CreateBuildWithManagedContainerRequestModel struct {
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container to create a build from.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // GameAssetReferences the list of game assets related to the build.
    GameAssetReferences []AssetReferenceParamsModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceParamsModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata to tag the build. The keys are case insensitive. The build metadata is made available to the server through
// Game Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports to map the build on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configurations for the build.
    RegionConfigurations []BuildRegionParamsModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server is started, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size to create the build on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateBuildWithManagedContainerResponse 
type CreateBuildWithManagedContainerResponseModel struct {
    // BuildId the guid string build ID. Must be unique for every build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // ContainerFlavor the flavor of container of the build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata the metadata of the build.
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to host on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// CreateMatchmakingTicketRequest the client specifies the creator's attributes and optionally a list of other users to match with.
type CreateMatchmakingTicketRequestModel struct {
    // Creator the User who created this ticket.
    Creator MatchmakingPlayerModel `json:"Creator,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MembersToMatchWith a list of Entity Keys of other users to match with.
    MembersToMatchWith []EntityKeyModel `json:"MembersToMatchWith,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
}

// CreateMatchmakingTicketResult 
type CreateMatchmakingTicketResultModel struct {
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CreateRemoteUserRequest creates a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type CreateRemoteUserRequestModel struct {
    // BuildId the guid string build ID of to create the remote user for.
    BuildId string `json:"BuildId,omitempty"`
    // ExpirationTime the expiration time for the remote user created. Defaults to expiring in one day if not specified.
    ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
    // Region the region of virtual machine to create the remote user for.
    Region AzureRegion `json:"Region,omitempty"`
    // Username the username to create the remote user with.
    Username string `json:"Username,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// CreateRemoteUserResponse 
type CreateRemoteUserResponseModel struct {
    // ExpirationTime the expiration time for the remote user created.
    ExpirationTime time.Time `json:"ExpirationTime,omitempty"`
    // Password the generated password for the remote user that was created.
    Password string `json:"Password,omitempty"`
    // Username the username for the remote user that was created.
    Username string `json:"Username,omitempty"`
}

// CreateServerBackfillTicketRequest the server specifies all the members, their teams and their attributes, and the server details if applicable.
type CreateServerBackfillTicketRequestModel struct {
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // Members the users who will be part of this ticket, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // ServerDetails the details of the server the members are connected to.
    ServerDetails *ServerDetailsModel `json:"ServerDetails,omitempty"`
}

// CreateServerBackfillTicketResult 
type CreateServerBackfillTicketResultModel struct {
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// CreateServerMatchmakingTicketRequest the server specifies all the members and their attributes.
type CreateServerMatchmakingTicketRequestModel struct {
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // Members the users who will be part of this ticket.
    Members []MatchmakingPlayerModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
}

// CurrentServerStats 
type CurrentServerStatsModel struct {
    // Active the number of active multiplayer servers.
    Active int32 `json:"Active,omitempty"`
    // Propping the number of multiplayer servers still downloading game resources (such as assets).
    Propping int32 `json:"Propping,omitempty"`
    // StandingBy the number of standingby multiplayer servers.
    StandingBy int32 `json:"StandingBy,omitempty"`
    // Total the total number of multiplayer servers.
    Total int32 `json:"Total,omitempty"`
}

// CustomDifferenceRuleExpansion 
type CustomDifferenceRuleExpansionModel struct {
    // DifferenceOverrides manually specify the values to use for each expansion interval (this overrides Difference, Delta, and MaxDifference).
    DifferenceOverrides []OverrideDoubleModel `json:"DifferenceOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// CustomRegionSelectionRuleExpansion 
type CustomRegionSelectionRuleExpansionModel struct {
    // MaxLatencyOverrides manually specify the maximum latency to use for each expansion interval.
    MaxLatencyOverrides []OverrideUnsignedIntModel `json:"MaxLatencyOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// CustomSetIntersectionRuleExpansion 
type CustomSetIntersectionRuleExpansionModel struct {
    // MinIntersectionSizeOverrides manually specify the values to use for each expansion interval.
    MinIntersectionSizeOverrides []OverrideUnsignedIntModel `json:"MinIntersectionSizeOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// CustomTeamDifferenceRuleExpansion 
type CustomTeamDifferenceRuleExpansionModel struct {
    // DifferenceOverrides manually specify the team difference value to use for each expansion interval.
    DifferenceOverrides []OverrideDoubleModel `json:"DifferenceOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// CustomTeamSizeBalanceRuleExpansion 
type CustomTeamSizeBalanceRuleExpansionModel struct {
    // DifferenceOverrides manually specify the team size difference to use for each expansion interval.
    DifferenceOverrides []OverrideUnsignedIntModel `json:"DifferenceOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// DeleteAssetRequest deletes a multiplayer server game asset for a title.
type DeleteAssetRequestModel struct {
    // FileName the filename of the asset to delete.
    FileName string `json:"FileName,omitempty"`
}

// DeleteBuildRequest deletes a multiplayer server build.
type DeleteBuildRequestModel struct {
    // BuildId the guid string build ID of the build to delete.
    BuildId string `json:"BuildId,omitempty"`
}

// DeleteCertificateRequest deletes a multiplayer server game certificate.
type DeleteCertificateRequestModel struct {
    // Name the name of the certificate.
    Name string `json:"Name,omitempty"`
}

// DeleteRemoteUserRequest deletes a remote user to log on to a VM for a multiplayer server build in a specific region. Returns user credential
// information necessary to log on.
type DeleteRemoteUserRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server where the remote user is to delete.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server where the remote user is to delete.
    Region AzureRegion `json:"Region,omitempty"`
    // Username the username of the remote user to delete.
    Username string `json:"Username,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// DifferenceRule 
type DifferenceRuleModel struct {
    // Attribute description of the attribute used by this rule to match tickets.
    Attribute QueueRuleAttributeModel `json:"Attribute,omitempty"`
    // AttributeNotSpecifiedBehavior describes the behavior when an attribute is not specified in the ticket creation request or in the user's entity
// profile.
    AttributeNotSpecifiedBehavior AttributeNotSpecifiedBehavior `json:"AttributeNotSpecifiedBehavior,omitempty"`
    // CustomExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule. When this
// is set, Difference is ignored.
    CustomExpansion *CustomDifferenceRuleExpansionModel `json:"CustomExpansion,omitempty"`
    // DefaultAttributeValue the default value assigned to tickets that are missing the attribute specified by AttributePath (assuming that
// AttributeNotSpecifiedBehavior is false). Optional.
    DefaultAttributeValue float64 `json:"DefaultAttributeValue,omitempty"`
    // Difference the allowed difference between any two tickets at the start of matchmaking.
    Difference float64 `json:"Difference,omitempty"`
    // LinearExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule.
    LinearExpansion *LinearDifferenceRuleExpansionModel `json:"LinearExpansion,omitempty"`
    // MergeFunction how values are treated when there are multiple players in a single ticket.
    MergeFunction AttributeMergeFunction `json:"MergeFunction,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
    // Weight the relative weight of this rule compared to others.
    Weight float64 `json:"Weight,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EnableMultiplayerServersForTitleRequest enables the multiplayer server feature for a title and returns the enabled status. The enabled status can be
// Initializing, Enabled, and Disabled. It can up to 20 minutes or more for the title to be enabled for the feature. On
// average, it can take up to 20 minutes for the title to be enabled for the feature.
type EnableMultiplayerServersForTitleRequestModel struct {
}

// EnableMultiplayerServersForTitleResponse 
type EnableMultiplayerServersForTitleResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// EnablePartiesForTitleRequest enables the parties feature for a title.
type EnablePartiesForTitleRequestModel struct {
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// GameCertificateReference 
type GameCertificateReferenceModel struct {
    // GsdkAlias an alias for the game certificate. The game server will reference this alias via GSDK config to retrieve the game
// certificate. This alias is used as an identifier in game server code to allow a new certificate with different Name
// field to be uploaded without the need to change any game server code to reference the new Name.
    GsdkAlias string `json:"GsdkAlias,omitempty"`
    // Name the name of the game certificate. This name should match the name of a certificate that was previously uploaded to this
// title.
    Name string `json:"Name,omitempty"`
}

// GameCertificateReferenceParams 
type GameCertificateReferenceParamsModel struct {
    // GsdkAlias an alias for the game certificate. The game server will reference this alias via GSDK config to retrieve the game
// certificate. This alias is used as an identifier in game server code to allow a new certificate with different Name
// field to be uploaded without the need to change any game server code to reference the new Name.
    GsdkAlias string `json:"GsdkAlias,omitempty"`
    // Name the name of the game certificate. This name should match the name of a certificate that was previously uploaded to this
// title.
    Name string `json:"Name,omitempty"`
}

// GetAssetUploadUrlRequest gets the URL to upload assets to.
type GetAssetUploadUrlRequestModel struct {
    // FileName the asset's file name to get the upload URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetAssetUploadUrlResponse 
type GetAssetUploadUrlResponseModel struct {
    // AssetUploadUrl the asset's upload URL.
    AssetUploadUrl string `json:"AssetUploadUrl,omitempty"`
    // FileName the asset's file name to get the upload URL for.
    FileName string `json:"FileName,omitempty"`
}

// GetBuildRequest returns the details about a multiplayer server build.
type GetBuildRequestModel struct {
    // BuildId the guid string build ID of the build to get.
    BuildId string `json:"BuildId,omitempty"`
}

// GetBuildResponse 
type GetBuildResponseModel struct {
    // BuildId the guid string build ID of the build.
    BuildId string `json:"BuildId,omitempty"`
    // BuildName the build name.
    BuildName string `json:"BuildName,omitempty"`
    // BuildStatus the current build status. Valid values are - Deploying, Deployed, DeletingRegion, Unhealthy.
    BuildStatus string `json:"BuildStatus,omitempty"`
    // ContainerFlavor the flavor of container of he build.
    ContainerFlavor ContainerFlavor `json:"ContainerFlavor,omitempty"`
    // ContainerRunCommand the container command to run when the multiplayer server has been allocated, including any arguments. This only applies
// to custom builds. If the build is a managed build, this field will be null.
    ContainerRunCommand string `json:"ContainerRunCommand,omitempty"`
    // CreationTime the time the build was created in UTC.
    CreationTime time.Time `json:"CreationTime,omitempty"`
    // CustomGameContainerImage the custom game container image for a custom build.
    CustomGameContainerImage *ContainerImageReferenceModel `json:"CustomGameContainerImage,omitempty"`
    // GameAssetReferences the game assets for the build.
    GameAssetReferences []AssetReferenceModel `json:"GameAssetReferences,omitempty"`
    // GameCertificateReferences the game certificates for the build.
    GameCertificateReferences []GameCertificateReferenceModel `json:"GameCertificateReferences,omitempty"`
    // Metadata metadata of the build. The keys are case insensitive. The build metadata is made available to the server through Game
// Server SDK (GSDK).
    Metadata map[string]string `json:"Metadata,omitempty"`
    // MultiplayerServerCountPerVm the number of multiplayer servers to hosted on a single VM of the build.
    MultiplayerServerCountPerVm int32 `json:"MultiplayerServerCountPerVm,omitempty"`
    // Ports the ports the build is mapped on.
    Ports []PortModel `json:"Ports,omitempty"`
    // RegionConfigurations the region configuration for the build.
    RegionConfigurations []BuildRegionModel `json:"RegionConfigurations,omitempty"`
    // StartMultiplayerServerCommand the command to run when the multiplayer server has been allocated, including any arguments. This only applies to managed
// builds. If the build is a custom build, this field will be null.
    StartMultiplayerServerCommand string `json:"StartMultiplayerServerCommand,omitempty"`
    // VmSize the VM size the build was created on.
    VmSize AzureVmSize `json:"VmSize,omitempty"`
}

// GetCognitiveServicesTokenRequest gets a token for the cognitive services where translation queries can be sent to.
type GetCognitiveServicesTokenRequestModel struct {
    // CognitiveServicesType the type of the cognitive service for which a token is being requested.
    CognitiveServicesType CognitiveServicesType `json:"CognitiveServicesType,omitempty"`
    // Region the region the client is closest to (A service endpoint will be provided for this region or one nearest one available).
    Region AzureRegion `json:"Region,omitempty"`
}

// GetCognitiveServicesTokenResponse 
type GetCognitiveServicesTokenResponseModel struct {
    // ServiceEndpoint the url for the service where further requests can be made with the given token.
    ServiceEndpoint string `json:"ServiceEndpoint,omitempty"`
    // Token the token for the cognitive service.
    Token string `json:"Token,omitempty"`
}

// GetContainerRegistryCredentialsRequest gets credentials to the container registry where game developers can upload custom container images to before creating a
// new build.
type GetContainerRegistryCredentialsRequestModel struct {
}

// GetContainerRegistryCredentialsResponse 
type GetContainerRegistryCredentialsResponseModel struct {
    // DnsName the url of the container registry.
    DnsName string `json:"DnsName,omitempty"`
    // Password the password for accessing the container registry.
    Password string `json:"Password,omitempty"`
    // Username the username for accessing the container registry.
    Username string `json:"Username,omitempty"`
}

// GetMatchmakingQueueRequest gets the current configuration for a queue.
type GetMatchmakingQueueRequestModel struct {
    // QueueName the Id of the matchmaking queue to retrieve.
    QueueName string `json:"QueueName,omitempty"`
}

// GetMatchmakingQueueResult 
type GetMatchmakingQueueResultModel struct {
    // MatchmakingQueue the matchmaking queue config.
    MatchmakingQueue *MatchmakingQueueConfigModel `json:"MatchmakingQueue,omitempty"`
}

// GetMatchmakingTicketRequest the ticket includes the invited players, their attributes if they have joined, the ticket status, the match Id when
// applicable, etc. Only servers, the ticket creator and the invited players can get the ticket.
type GetMatchmakingTicketRequestModel struct {
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject,omitempty"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetMatchmakingTicketResult 
type GetMatchmakingTicketResultModel struct {
    // CancellationReason the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReason CancellationReason `json:"CancellationReason,omitempty"`
    // CancellationReasonString the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReasonString string `json:"CancellationReasonString,omitempty"`
    // Created the server date and time at which ticket was created.
    Created time.Time `json:"Created,omitempty"`
    // Creator the Creator's entity key.
    Creator EntityKeyModel `json:"Creator,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that have joined this ticket.
    Members []MatchmakingPlayerModel `json:"Members,omitempty"`
    // MembersToMatchWith a list of PlayFab Ids of Users to match with.
    MembersToMatchWith []EntityKeyModel `json:"MembersToMatchWith,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // Status the current ticket status. Possible values are: WaitingForPlayers, WaitingForMatch, WaitingForServer, Canceled and
// Matched.
    Status string `json:"Status,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetMatchRequest when matchmaking has successfully matched together a collection of tickets, it produces a 'match' with an Id. The match
// contains all of the players that were matched together, and their team assigments. Only servers and ticket members can
// get the match.
type GetMatchRequestModel struct {
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // QueueName the name of the queue to join.
    QueueName string `json:"QueueName,omitempty"`
    // ReturnMemberAttributes determines whether the matchmaking attributes for each user should be returned in the response for match request.
    ReturnMemberAttributes bool `json:"ReturnMemberAttributes,omitempty"`
}

// GetMatchResult 
type GetMatchResultModel struct {
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that are matched together, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // RegionPreferences a list of regions that the match could be played in sorted by preference. This value is only set if the queue has a
// region selection rule.
    RegionPreferences []string `json:"RegionPreferences,omitempty"`
    // ServerDetails the details of the server that the match has been allocated to.
    ServerDetails *ServerDetailsModel `json:"ServerDetails,omitempty"`
}

// GetMultiplayerServerDetailsRequest gets multiplayer server session details for a build in a specific region.
type GetMultiplayerServerDetailsRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to get details for.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region the multiplayer server is located in to get details for.
    Region AzureRegion `json:"Region,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server to get details for. This is to keep track of
// multiplayer server sessions.
    SessionId string `json:"SessionId,omitempty"`
}

// GetMultiplayerServerDetailsResponse 
type GetMultiplayerServerDetailsResponseModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // FQDN the fully qualified domain name of the virtual machine that is hosting this multiplayer server.
    FQDN string `json:"FQDN,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// GetMultiplayerServerLogsRequest gets multiplayer server logs for a specific server id in a region. The logs are available only after a server has
// terminated.
type GetMultiplayerServerLogsRequestModel struct {
    // Region the region of the multiplayer server to get logs for.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the server ID of multiplayer server to get logs for.
    ServerId string `json:"ServerId,omitempty"`
}

// GetMultiplayerServerLogsResponse 
type GetMultiplayerServerLogsResponseModel struct {
    // LogDownloadUrl 
    LogDownloadUrl string `json:"LogDownloadUrl,omitempty"`
}

// GetQueueStatisticsRequest returns the matchmaking statistics for a queue. These include the number of players matching and the statistics related
// to the time to match statistics in seconds (average and percentiles). Statistics are refreshed once every 5 minutes.
// Servers can access all statistics no matter what the ClientStatisticsVisibility is configured to. Clients can access
// statistics according to the ClientStatisticsVisibility. Client requests are forbidden if all visibility fields are
// false.
type GetQueueStatisticsRequestModel struct {
    // QueueName the name of the queue.
    QueueName string `json:"QueueName,omitempty"`
}

// GetQueueStatisticsResult 
type GetQueueStatisticsResultModel struct {
    // NumberOfPlayersMatching the current number of players in the matchmaking queue, who are waiting to be matched.
    NumberOfPlayersMatching uint32 `json:"NumberOfPlayersMatching,omitempty"`
    // TimeToMatchStatisticsInSeconds statistics representing the time (in seconds) it takes for tickets to find a match.
    TimeToMatchStatisticsInSeconds *StatisticsModel `json:"TimeToMatchStatisticsInSeconds,omitempty"`
}

// GetRemoteLoginEndpointRequest gets a remote login endpoint to a VM that is hosting a multiplayer server build in a specific region.
type GetRemoteLoginEndpointRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to get remote login information for.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server to get remote login information for.
    Region AzureRegion `json:"Region,omitempty"`
    // VmId the virtual machine ID the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// GetRemoteLoginEndpointResponse 
type GetRemoteLoginEndpointResponseModel struct {
    // IPV4Address the remote login IPV4 address of multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // Port the remote login port of multiplayer server.
    Port int32 `json:"Port,omitempty"`
}

// GetServerBackfillTicketRequest the ticket includes the players, their attributes, their teams, the ticket status, the match Id and the server details
// when applicable, etc. Only servers can get the ticket.
type GetServerBackfillTicketRequestModel struct {
    // EscapeObject determines whether the matchmaking attributes will be returned as an escaped JSON string or as an un-escaped JSON
// object.
    EscapeObject bool `json:"EscapeObject,omitempty"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetServerBackfillTicketResult 
type GetServerBackfillTicketResultModel struct {
    // CancellationReason the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReason CancellationReason `json:"CancellationReason,omitempty"`
    // CancellationReasonString the reason why the current ticket was canceled. This field is only set if the ticket is in canceled state.
    CancellationReasonString string `json:"CancellationReasonString,omitempty"`
    // Created the server date and time at which ticket was created.
    Created time.Time `json:"Created,omitempty"`
    // GiveUpAfterSeconds how long to attempt matching this ticket in seconds.
    GiveUpAfterSeconds int32 `json:"GiveUpAfterSeconds,omitempty"`
    // MatchId the Id of a match.
    MatchId string `json:"MatchId,omitempty"`
    // Members a list of Users that are part of this ticket, along with their team assignments.
    Members []MatchmakingPlayerWithTeamAssignmentModel `json:"Members,omitempty"`
    // QueueName the Id of a match queue.
    QueueName string `json:"QueueName,omitempty"`
    // ServerDetails the details of the server the members are connected to.
    ServerDetails ServerDetailsModel `json:"ServerDetails,omitempty"`
    // Status the current ticket status. Possible values are: WaitingForMatch, Canceled and Matched.
    Status string `json:"Status,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// GetTitleEnabledForMultiplayerServersStatusRequest gets the status of whether a title is enabled for the multiplayer server feature. The enabled status can be
// Initializing, Enabled, and Disabled.
type GetTitleEnabledForMultiplayerServersStatusRequestModel struct {
}

// GetTitleEnabledForMultiplayerServersStatusResponse 
type GetTitleEnabledForMultiplayerServersStatusResponseModel struct {
    // Status the enabled status for the multiplayer server features for the title.
    Status TitleMultiplayerServerEnabledStatus `json:"Status,omitempty"`
}

// GetTitleMultiplayerServersQuotasRequest gets the quotas for a title in relation to multiplayer servers.
type GetTitleMultiplayerServersQuotasRequestModel struct {
}

// GetTitleMultiplayerServersQuotasResponse 
type GetTitleMultiplayerServersQuotasResponseModel struct {
    // Quotas the various quotas for multiplayer servers for the title.
    Quotas *TitleMultiplayerServersQuotasModel `json:"Quotas,omitempty"`
}

// JoinMatchmakingTicketRequest add the player to a matchmaking ticket and specify all of its matchmaking attributes. Players can join a ticket if and
// only if their EntityKeys are already listed in the ticket's Members list. The matchmaking service automatically starts
// matching the ticket against other matchmaking tickets once all players have joined the ticket. It is not possible to
// join a ticket once it has started matching.
type JoinMatchmakingTicketRequestModel struct {
    // Member the User who wants to join the ticket. Their Id must be listed in PlayFabIdsToMatchWith.
    Member MatchmakingPlayerModel `json:"Member,omitempty"`
    // QueueName the name of the queue to join.
    QueueName string `json:"QueueName,omitempty"`
    // TicketId the Id of the ticket to find a match for.
    TicketId string `json:"TicketId,omitempty"`
}

// JoinMatchmakingTicketResult 
type JoinMatchmakingTicketResultModel struct {
}

// LinearDifferenceRuleExpansion 
type LinearDifferenceRuleExpansionModel struct {
    // Delta this value gets added to Difference at every expansion interval.
    Delta float64 `json:"Delta,omitempty"`
    // Limit once the total difference reaches this value, expansion stops. Optional.
    Limit float64 `json:"Limit,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// LinearRegionSelectionRuleExpansion 
type LinearRegionSelectionRuleExpansionModel struct {
    // Delta this value gets added to MaxLatency at every expansion interval.
    Delta uint32 `json:"Delta,omitempty"`
    // Limit once the max Latency reaches this value, expansion stops.
    Limit uint32 `json:"Limit,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// LinearSetIntersectionRuleExpansion 
type LinearSetIntersectionRuleExpansionModel struct {
    // Delta this value gets added to MinIntersectionSize at every expansion interval.
    Delta uint32 `json:"Delta,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// LinearTeamDifferenceRuleExpansion 
type LinearTeamDifferenceRuleExpansionModel struct {
    // Delta this value gets added to Difference at every expansion interval.
    Delta float64 `json:"Delta,omitempty"`
    // Limit once the total difference reaches this value, expansion stops. Optional.
    Limit float64 `json:"Limit,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// LinearTeamSizeBalanceRuleExpansion 
type LinearTeamSizeBalanceRuleExpansionModel struct {
    // Delta this value gets added to Difference at every expansion interval.
    Delta uint32 `json:"Delta,omitempty"`
    // Limit once the total difference reaches this value, expansion stops. Optional.
    Limit uint32 `json:"Limit,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// ListAssetSummariesRequest returns a list of multiplayer server game asset summaries for a title.
type ListAssetSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListAssetSummariesResponse 
type ListAssetSummariesResponseModel struct {
    // AssetSummaries the list of asset summaries.
    AssetSummaries []AssetSummaryModel `json:"AssetSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildSummariesRequest returns a list of summarized details of all multiplayer server builds for a title.
type ListBuildSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListBuildSummariesResponse 
type ListBuildSummariesResponseModel struct {
    // BuildSummaries the list of build summaries for a title.
    BuildSummaries []BuildSummaryModel `json:"BuildSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListCertificateSummariesRequest returns a list of multiplayer server game certificates for a title.
type ListCertificateSummariesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListCertificateSummariesResponse 
type ListCertificateSummariesResponseModel struct {
    // CertificateSummaries the list of game certificates.
    CertificateSummaries []CertificateSummaryModel `json:"CertificateSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImagesRequest returns a list of the container images that have been uploaded to the container registry for a title.
type ListContainerImagesRequestModel struct {
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImagesResponse 
type ListContainerImagesResponseModel struct {
    // Images the list of container images.
    Images []string `json:"Images,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListContainerImageTagsRequest returns a list of the tags for a particular container image that exists in the container registry for a title.
type ListContainerImageTagsRequestModel struct {
    // ImageName the container images we want to list tags for.
    ImageName string `json:"ImageName,omitempty"`
}

// ListContainerImageTagsResponse 
type ListContainerImageTagsResponseModel struct {
    // Tags the list of tags for a particular container image.
    Tags []string `json:"Tags,omitempty"`
}

// ListMatchmakingQueuesRequest gets a list of all the matchmaking queue configurations for the title.
type ListMatchmakingQueuesRequestModel struct {
}

// ListMatchmakingQueuesResult 
type ListMatchmakingQueuesResultModel struct {
    // MatchMakingQueues the list of matchmaking queue configs for this title.
    MatchMakingQueues []MatchmakingQueueConfigModel `json:"MatchMakingQueues,omitempty"`
}

// ListMatchmakingTicketsForPlayerRequest if the caller is a title, the EntityKey in the request is required. If the caller is a player, then it is optional. If
// it is provided it must match the caller's entity.
type ListMatchmakingTicketsForPlayerRequestModel struct {
    // Entity the entity key for which to find the ticket Ids.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue to find a match for.
    QueueName string `json:"QueueName,omitempty"`
}

// ListMatchmakingTicketsForPlayerResult 
type ListMatchmakingTicketsForPlayerResultModel struct {
    // TicketIds the list of ticket Ids the user is a member of.
    TicketIds []string `json:"TicketIds,omitempty"`
}

// ListMultiplayerServersRequest returns a list of multiplayer servers for a build in a specific region.
type ListMultiplayerServersRequestModel struct {
    // BuildId the guid string build ID of the multiplayer servers to list.
    BuildId string `json:"BuildId,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region the multiplayer servers to list.
    Region AzureRegion `json:"Region,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListMultiplayerServersResponse 
type ListMultiplayerServersResponseModel struct {
    // MultiplayerServerSummaries the list of multiplayer server summary details.
    MultiplayerServerSummaries []MultiplayerServerSummaryModel `json:"MultiplayerServerSummaries,omitempty"`
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListQosServersForTitleRequest returns a list of quality of service servers for a title.
type ListQosServersForTitleRequestModel struct {
}

// ListQosServersForTitleResponse 
type ListQosServersForTitleResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // QosServers the list of QoS servers.
    QosServers []QosServerModel `json:"QosServers,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListQosServersRequest returns a list of quality of service servers.
type ListQosServersRequestModel struct {
}

// ListQosServersResponse 
type ListQosServersResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // QosServers the list of QoS servers.
    QosServers []QosServerModel `json:"QosServers,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListServerBackfillTicketsForPlayerRequest list all server backfill ticket Ids the user is a member of.
type ListServerBackfillTicketsForPlayerRequestModel struct {
    // Entity the entity key for which to find the ticket Ids.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // QueueName the name of the queue the tickets are in.
    QueueName string `json:"QueueName,omitempty"`
}

// ListServerBackfillTicketsForPlayerResult 
type ListServerBackfillTicketsForPlayerResultModel struct {
    // TicketIds the list of backfill ticket Ids the user is a member of.
    TicketIds []string `json:"TicketIds,omitempty"`
}

// ListVirtualMachineSummariesRequest returns a list of virtual machines for a title.
type ListVirtualMachineSummariesRequestModel struct {
    // BuildId the guid string build ID of the virtual machines to list.
    BuildId string `json:"BuildId,omitempty"`
    // PageSize the page size for the request.
    PageSize int32 `json:"PageSize,omitempty"`
    // Region the region of the virtual machines to list.
    Region AzureRegion `json:"Region,omitempty"`
    // SkipToken the skip token for the paged request.
    SkipToken string `json:"SkipToken,omitempty"`
}

// ListVirtualMachineSummariesResponse 
type ListVirtualMachineSummariesResponseModel struct {
    // PageSize the page size on the response.
    PageSize int32 `json:"PageSize,omitempty"`
    // SkipToken the skip token for the paged response.
    SkipToken string `json:"SkipToken,omitempty"`
    // VirtualMachines the list of virtual machine summaries.
    VirtualMachines []VirtualMachineSummaryModel `json:"VirtualMachines,omitempty"`
}

// MatchmakingPlayer a user in a matchmaking ticket.
type MatchmakingPlayerModel struct {
    // Attributes the user's attributes custom to the title.
    Attributes *MatchmakingPlayerAttributesModel `json:"Attributes,omitempty"`
    // Entity the entity key of the matchmaking user.
    Entity EntityKeyModel `json:"Entity,omitempty"`
}

// MatchmakingPlayerAttributes the matchmaking attributes for a user.
type MatchmakingPlayerAttributesModel struct {
    // DataObject a data object representing a user's attributes.
    DataObject interface{} `json:"DataObject,omitempty"`
    // EscapedDataObject an escaped data object representing a user's attributes.
    EscapedDataObject string `json:"EscapedDataObject,omitempty"`
}

// MatchmakingPlayerWithTeamAssignment a player in a created matchmaking match with a team assignment.
type MatchmakingPlayerWithTeamAssignmentModel struct {
    // Attributes the user's attributes custom to the title. These attributes will be null unless the request has ReturnMemberAttributes
// flag set to true.
    Attributes *MatchmakingPlayerAttributesModel `json:"Attributes,omitempty"`
    // Entity the entity key of the matchmaking user.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // TeamId the Id of the team the User is assigned to.
    TeamId string `json:"TeamId,omitempty"`
}

// MatchmakingQueueConfig 
type MatchmakingQueueConfigModel struct {
    // BuildId this is the buildId that will be used to allocate the multiplayer server for the match.
    BuildId string `json:"BuildId,omitempty"`
    // DifferenceRules list of difference rules used to find an optimal match.
    DifferenceRules []DifferenceRuleModel `json:"DifferenceRules,omitempty"`
    // MatchTotalRules list of match total rules used to find an optimal match.
    MatchTotalRules []MatchTotalRuleModel `json:"MatchTotalRules,omitempty"`
    // MaxMatchSize maximum number of players in a match.
    MaxMatchSize uint32 `json:"MaxMatchSize,omitempty"`
    // MaxTicketSize maximum number of players in a ticket. Optional.
    MaxTicketSize uint32 `json:"MaxTicketSize,omitempty"`
    // MinMatchSize minimum number of players in a match.
    MinMatchSize uint32 `json:"MinMatchSize,omitempty"`
    // Name unique identifier for a Queue. Chosen by the developer.
    Name string `json:"Name,omitempty"`
    // RegionSelectionRule region selection rule used to find an optimal match.
    RegionSelectionRule *RegionSelectionRuleModel `json:"RegionSelectionRule,omitempty"`
    // ServerAllocationEnabled boolean flag to enable server allocation for the queue.
    ServerAllocationEnabled bool `json:"ServerAllocationEnabled,omitempty"`
    // SetIntersectionRules list of set intersection rules used to find an optimal match.
    SetIntersectionRules []SetIntersectionRuleModel `json:"SetIntersectionRules,omitempty"`
    // StatisticsVisibilityToPlayers controls which statistics are visible to players.
    StatisticsVisibilityToPlayers *StatisticsVisibilityToPlayersModel `json:"StatisticsVisibilityToPlayers,omitempty"`
    // StringEqualityRules list of string equality rules used to find an optimal match.
    StringEqualityRules []StringEqualityRuleModel `json:"StringEqualityRules,omitempty"`
    // TeamDifferenceRules list of team difference rules used to find an optimal match.
    TeamDifferenceRules []TeamDifferenceRuleModel `json:"TeamDifferenceRules,omitempty"`
    // Teams the team configuration for a match. This may be null if there are no teams.
    Teams []MatchmakingQueueTeamModel `json:"Teams,omitempty"`
    // TeamSizeBalanceRule team size balance rule used to find an optimal match.
    TeamSizeBalanceRule *TeamSizeBalanceRuleModel `json:"TeamSizeBalanceRule,omitempty"`
    // TeamTicketSizeSimilarityRule team ticket size similarity rule used to find an optimal match.
    TeamTicketSizeSimilarityRule *TeamTicketSizeSimilarityRuleModel `json:"TeamTicketSizeSimilarityRule,omitempty"`
}

// MatchmakingQueueTeam 
type MatchmakingQueueTeamModel struct {
    // MaxTeamSize the maximum number of players required for the team.
    MaxTeamSize uint32 `json:"MaxTeamSize,omitempty"`
    // MinTeamSize the minimum number of players required for the team.
    MinTeamSize uint32 `json:"MinTeamSize,omitempty"`
    // Name a name to identify the team. This is case insensitive.
    Name string `json:"Name,omitempty"`
}

// MatchTotalRule 
type MatchTotalRuleModel struct {
    // Attribute description of the attribute used by this rule to match tickets.
    Attribute QueueRuleAttributeModel `json:"Attribute,omitempty"`
    // Expansion collection of fields relating to expanding this rule at set intervals.
    Expansion *MatchTotalRuleExpansionModel `json:"Expansion,omitempty"`
    // Max the maximum total value for a group. Must be >= Min.
    Max float64 `json:"Max,omitempty"`
    // Min the minimum total value for a group. Must be >=2.
    Min float64 `json:"Min,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
    // Weight the relative weight of this rule compared to others.
    Weight float64 `json:"Weight,omitempty"`
}

// MatchTotalRuleExpansion 
type MatchTotalRuleExpansionModel struct {
    // MaxOverrides manually specify the values to use for each expansion interval. When this is set, Max is ignored.
    MaxOverrides []OverrideDoubleModel `json:"MaxOverrides,omitempty"`
    // MinOverrides manually specify the values to use for each expansion interval. When this is set, Min is ignored.
    MinOverrides []OverrideDoubleModel `json:"MinOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// MultiplayerServerSummary 
type MultiplayerServerSummaryModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the title generated guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// OverrideDouble 
type OverrideDoubleModel struct {
    // Value the custom expansion value.
    Value float64 `json:"Value,omitempty"`
}

// OverrideUnsignedInt 
type OverrideUnsignedIntModel struct {
    // Value the custom expansion value.
    Value uint32 `json:"Value,omitempty"`
}

// Port 
type PortModel struct {
    // Name the name for the port.
    Name string `json:"Name,omitempty"`
    // Num the number for the port.
    Num int32 `json:"Num,omitempty"`
    // Protocol the protocol for the port.
    Protocol ProtocolType `json:"Protocol,omitempty"`
}

// ProtocolType 
type ProtocolType string
  
const (
     ProtocolTypeTCP ProtocolType = "TCP"
     ProtocolTypeUDP ProtocolType = "UDP"
      )
// QosServer 
type QosServerModel struct {
    // Region the region the QoS server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerUrl the QoS server URL.
    ServerUrl string `json:"ServerUrl,omitempty"`
}

// QueueRuleAttribute 
type QueueRuleAttributeModel struct {
    // Path specifies which attribute in a ticket to use.
    Path string `json:"Path,omitempty"`
    // Source specifies which source the attribute comes from.
    Source AttributeSource `json:"Source,omitempty"`
}

// RegionSelectionRule 
type RegionSelectionRuleModel struct {
    // CustomExpansion controls how the Max Latency parameter expands over time. Only one expansion can be set per rule. When this is set,
// MaxLatency is ignored.
    CustomExpansion *CustomRegionSelectionRuleExpansionModel `json:"CustomExpansion,omitempty"`
    // LinearExpansion controls how the Max Latency parameter expands over time. Only one expansion can be set per rule.
    LinearExpansion *LinearRegionSelectionRuleExpansionModel `json:"LinearExpansion,omitempty"`
    // MaxLatency specifies the maximum latency that is allowed between the client and the selected server. The value is in milliseconds.
    MaxLatency uint32 `json:"MaxLatency,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // Path specifies which attribute in a ticket to use.
    Path string `json:"Path,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
    // Weight the relative weight of this rule compared to others.
    Weight float64 `json:"Weight,omitempty"`
}

// RemoveMatchmakingQueueRequest deletes the configuration for a queue. This will permanently delete the configuration and players will no longer be able
// to match in the queue. All outstanding matchmaking tickets will be cancelled.
type RemoveMatchmakingQueueRequestModel struct {
    // QueueName the Id of the matchmaking queue to remove.
    QueueName string `json:"QueueName,omitempty"`
}

// RemoveMatchmakingQueueResult 
type RemoveMatchmakingQueueResultModel struct {
}

// RequestMultiplayerServerRequest requests a multiplayer server session from a particular build in any of the given preferred regions.
type RequestMultiplayerServerRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to request.
    BuildId string `json:"BuildId,omitempty"`
    // InitialPlayers initial list of players (potentially matchmade) allowed to connect to the game. This list is passed to the game server
// when requested (via GSDK) and can be used to validate players connecting to it.
    InitialPlayers []string `json:"InitialPlayers,omitempty"`
    // PreferredRegions the preferred regions to request a multiplayer server from. The Multiplayer Service will iterate through the regions in
// the specified order and allocate a server from the first one that has servers available.
    PreferredRegions []AzureRegion `json:"PreferredRegions,omitempty"`
    // SessionCookie data encoded as a string that is passed to the game server when requested. This can be used to to communicate
// information such as game mode or map through the request flow.
    SessionCookie string `json:"SessionCookie,omitempty"`
    // SessionId a guid string session ID created track the multiplayer server session over its life.
    SessionId string `json:"SessionId,omitempty"`
}

// RequestMultiplayerServerResponse 
type RequestMultiplayerServerResponseModel struct {
    // ConnectedPlayers the connected players in the multiplayer server.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // FQDN the fully qualified domain name of the virtual machine that is hosting this multiplayer server.
    FQDN string `json:"FQDN,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the multiplayer server state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the region the multiplayer server is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the multiplayer server generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // SessionId the guid string session ID of the multiplayer server.
    SessionId string `json:"SessionId,omitempty"`
    // State the state of the multiplayer server.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the multiplayer server is located on.
    VmId string `json:"VmId,omitempty"`
}

// RequestPartyRequest requests a party session from a particular build in any of the given preferred regions.
type RequestPartyRequestModel struct {
    // PartyId a guid string party ID created track the party session over its life.
    PartyId string `json:"PartyId,omitempty"`
    // PreferredRegions the preferred regions to request a party session from. The party service will iterate through the regions in the
// specified order and allocate a party session from the first one that is available.
    PreferredRegions []AzureRegion `json:"PreferredRegions,omitempty"`
    // SessionCookie data encoded as a string that is passed to the party when requested. This can be used to to communicate information such
// as party type through the request flow.
    SessionCookie string `json:"SessionCookie,omitempty"`
    // Version the client version for the party being requested.
    Version string `json:"Version,omitempty"`
}

// RequestPartyResponse 
type RequestPartyResponseModel struct {
    // ConnectedPlayers the connected players in the party.
    ConnectedPlayers []ConnectedPlayerModel `json:"ConnectedPlayers,omitempty"`
    // DTLSCertificateSHA2Thumbprint the thumbprint of the DTLS certificate presented by the party session
    DTLSCertificateSHA2Thumbprint string `json:"DTLSCertificateSHA2Thumbprint,omitempty"`
    // FQDN the fully qualified domain name of the virtual machine that is hosting this party session.
    FQDN string `json:"FQDN,omitempty"`
    // IPV4Address the IPv4 address of the virtual machine that is hosting this party session.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // LastStateTransitionTime the time (UTC) at which a change in the party state was observed.
    LastStateTransitionTime time.Time `json:"LastStateTransitionTime,omitempty"`
    // PartyId the guid string party ID of the party session.
    PartyId string `json:"PartyId,omitempty"`
    // Ports the ports the party session uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the region the party session is located in.
    Region AzureRegion `json:"Region,omitempty"`
    // ServerId the string server ID of the party session generated by PlayFab.
    ServerId string `json:"ServerId,omitempty"`
    // State the state of the party session.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID that the party session is located on.
    VmId string `json:"VmId,omitempty"`
}

// RolloverContainerRegistryCredentialsRequest gets new credentials to the container registry where game developers can upload custom container images to before
// creating a new build.
type RolloverContainerRegistryCredentialsRequestModel struct {
}

// RolloverContainerRegistryCredentialsResponse 
type RolloverContainerRegistryCredentialsResponseModel struct {
    // DnsName the url of the container registry.
    DnsName string `json:"DnsName,omitempty"`
    // Password the password for accessing the container registry.
    Password string `json:"Password,omitempty"`
    // Username the username for accessing the container registry.
    Username string `json:"Username,omitempty"`
}

// ServerDetails 
type ServerDetailsModel struct {
    // IPV4Address the IPv4 address of the virtual machine that is hosting this multiplayer server.
    IPV4Address string `json:"IPV4Address,omitempty"`
    // Ports the ports the multiplayer server uses.
    Ports []PortModel `json:"Ports,omitempty"`
    // Region the server's region.
    Region string `json:"Region,omitempty"`
}

// SetIntersectionRule 
type SetIntersectionRuleModel struct {
    // Attribute description of the attribute used by this rule to match tickets.
    Attribute QueueRuleAttributeModel `json:"Attribute,omitempty"`
    // AttributeNotSpecifiedBehavior describes the behavior when an attribute is not specified in the ticket creation request or in the user's entity
// profile.
    AttributeNotSpecifiedBehavior AttributeNotSpecifiedBehavior `json:"AttributeNotSpecifiedBehavior,omitempty"`
    // CustomExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule. When this
// is set, MinIntersectionSize is ignored.
    CustomExpansion *CustomSetIntersectionRuleExpansionModel `json:"CustomExpansion,omitempty"`
    // DefaultAttributeValue the default value assigned to tickets that are missing the attribute specified by AttributePath (assuming that
// AttributeNotSpecifiedBehavior is UseDefault). Values must be unique.
    DefaultAttributeValue []string `json:"DefaultAttributeValue,omitempty"`
    // LinearExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule.
    LinearExpansion *LinearSetIntersectionRuleExpansionModel `json:"LinearExpansion,omitempty"`
    // MinIntersectionSize the minimum number of values that must match between sets.
    MinIntersectionSize uint32 `json:"MinIntersectionSize,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
    // Weight the relative weight of this rule compared to others.
    Weight float64 `json:"Weight,omitempty"`
}

// SetMatchmakingQueueRequest use this API to create or update matchmaking queue configurations. The queue configuration defines the matchmaking
// rules. The matchmaking service will match tickets together according to the configured rules. Queue resources are not
// spun up by calling this API. Queues are created when the first ticket is submitted.
type SetMatchmakingQueueRequestModel struct {
    // MatchmakingQueue the matchmaking queue config.
    MatchmakingQueue *MatchmakingQueueConfigModel `json:"MatchmakingQueue,omitempty"`
}

// SetMatchmakingQueueResult 
type SetMatchmakingQueueResultModel struct {
}

// ShutdownMultiplayerServerRequest executes the shutdown callback from the GSDK and terminates the multiplayer server session. The callback in the GSDK
// will allow for graceful shutdown with a 15 minute timeoutIf graceful shutdown has not been completed before 15 minutes
// have elapsed, the multiplayer server session will be forcefully terminated on it's own.
type ShutdownMultiplayerServerRequestModel struct {
    // BuildId the guid string build ID of the multiplayer server to delete.
    BuildId string `json:"BuildId,omitempty"`
    // Region the region of the multiplayer server to shut down.
    Region AzureRegion `json:"Region,omitempty"`
    // SessionId a guid string session ID of the multiplayer server to shut down.
    SessionId string `json:"SessionId,omitempty"`
}

// Statistics 
type StatisticsModel struct {
    // Average the average.
    Average float64 `json:"Average,omitempty"`
    // Percentile50 the 50th percentile.
    Percentile50 float64 `json:"Percentile50,omitempty"`
    // Percentile90 the 90th percentile.
    Percentile90 float64 `json:"Percentile90,omitempty"`
    // Percentile99 the 99th percentile.
    Percentile99 float64 `json:"Percentile99,omitempty"`
}

// StatisticsVisibilityToPlayers 
type StatisticsVisibilityToPlayersModel struct {
    // ShowNumberOfPlayersMatching whether to allow players to view the current number of players in the matchmaking queue.
    ShowNumberOfPlayersMatching bool `json:"ShowNumberOfPlayersMatching,omitempty"`
    // ShowTimeToMatch whether to allow players to view statistics representing the time it takes for tickets to find a match.
    ShowTimeToMatch bool `json:"ShowTimeToMatch,omitempty"`
}

// StringEqualityRule 
type StringEqualityRuleModel struct {
    // Attribute description of the attribute used by this rule to match tickets.
    Attribute QueueRuleAttributeModel `json:"Attribute,omitempty"`
    // AttributeNotSpecifiedBehavior describes the behavior when an attribute is not specified in the ticket creation request or in the user's entity
// profile.
    AttributeNotSpecifiedBehavior AttributeNotSpecifiedBehavior `json:"AttributeNotSpecifiedBehavior,omitempty"`
    // DefaultAttributeValue the default value assigned to tickets that are missing the attribute specified by AttributePath (assuming that
// AttributeNotSpecifiedBehavior is false).
    DefaultAttributeValue string `json:"DefaultAttributeValue,omitempty"`
    // Expansion collection of fields relating to expanding this rule at set intervals. For StringEqualityRules, this is limited to
// turning the rule off or on during different intervals.
    Expansion *StringEqualityRuleExpansionModel `json:"Expansion,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
    // Weight the relative weight of this rule compared to others.
    Weight float64 `json:"Weight,omitempty"`
}

// StringEqualityRuleExpansion 
type StringEqualityRuleExpansionModel struct {
    // EnabledOverrides list of bools specifying whether the rule is applied during this expansion.
    EnabledOverrides []bool `json:"EnabledOverrides,omitempty"`
    // SecondsBetweenExpansions how many seconds before this rule is expanded.
    SecondsBetweenExpansions uint32 `json:"SecondsBetweenExpansions,omitempty"`
}

// TeamDifferenceRule 
type TeamDifferenceRuleModel struct {
    // Attribute description of the attribute used by this rule to match teams.
    Attribute QueueRuleAttributeModel `json:"Attribute,omitempty"`
    // CustomExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule. When this
// is set, Difference is ignored.
    CustomExpansion *CustomTeamDifferenceRuleExpansionModel `json:"CustomExpansion,omitempty"`
    // DefaultAttributeValue the default value assigned to tickets that are missing the attribute specified by AttributePath (assuming that
// AttributeNotSpecifiedBehavior is false).
    DefaultAttributeValue float64 `json:"DefaultAttributeValue,omitempty"`
    // Difference the allowed difference between any two teams at the start of matchmaking.
    Difference float64 `json:"Difference,omitempty"`
    // LinearExpansion collection of fields relating to expanding this rule at set intervals. Only one expansion can be set per rule.
    LinearExpansion *LinearTeamDifferenceRuleExpansionModel `json:"LinearExpansion,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
}

// TeamSizeBalanceRule 
type TeamSizeBalanceRuleModel struct {
    // CustomExpansion controls how the Difference parameter expands over time. Only one expansion can be set per rule. When this is set,
// Difference is ignored.
    CustomExpansion *CustomTeamSizeBalanceRuleExpansionModel `json:"CustomExpansion,omitempty"`
    // Difference the allowed difference in team size between any two teams.
    Difference uint32 `json:"Difference,omitempty"`
    // LinearExpansion controls how the Difference parameter expands over time. Only one expansion can be set per rule.
    LinearExpansion *LinearTeamSizeBalanceRuleExpansionModel `json:"LinearExpansion,omitempty"`
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
}

// TeamTicketSizeSimilarityRule 
type TeamTicketSizeSimilarityRuleModel struct {
    // Name friendly name chosen by developer.
    Name string `json:"Name,omitempty"`
    // SecondsUntilOptional how many seconds before this rule is no longer enforced (but tickets that comply with this rule will still be
// prioritized over those that don't). Leave blank if this rule is always enforced.
    SecondsUntilOptional uint32 `json:"SecondsUntilOptional,omitempty"`
}

// TitleMultiplayerServerEnabledStatus 
type TitleMultiplayerServerEnabledStatus string
  
const (
     TitleMultiplayerServerEnabledStatusInitializing TitleMultiplayerServerEnabledStatus = "Initializing"
     TitleMultiplayerServerEnabledStatusEnabled TitleMultiplayerServerEnabledStatus = "Enabled"
     TitleMultiplayerServerEnabledStatusDisabled TitleMultiplayerServerEnabledStatus = "Disabled"
      )
// TitleMultiplayerServersQuotas 
type TitleMultiplayerServersQuotasModel struct {
    // CoreCapacities the core capacity for the various regions and VM Family
    CoreCapacities []CoreCapacityModel `json:"CoreCapacities,omitempty"`
}

// UpdateBuildRegionsRequest updates a multiplayer server build's regions.
type UpdateBuildRegionsRequestModel struct {
    // BuildId the guid string ID of the build we want to update regions for.
    BuildId string `json:"BuildId,omitempty"`
    // BuildRegions the updated region configuration that should be applied to the specified build.
    BuildRegions []BuildRegionParamsModel `json:"BuildRegions,omitempty"`
}

// UploadCertificateRequest uploads a multiplayer server game certificate.
type UploadCertificateRequestModel struct {
    // GameCertificate the game certificate to upload.
    GameCertificate CertificateModel `json:"GameCertificate,omitempty"`
}

// VirtualMachineSummary 
type VirtualMachineSummaryModel struct {
    // HealthStatus the virtual machine health status.
    HealthStatus string `json:"HealthStatus,omitempty"`
    // State the virtual machine state.
    State string `json:"State,omitempty"`
    // VmId the virtual machine ID.
    VmId string `json:"VmId,omitempty"`
}
