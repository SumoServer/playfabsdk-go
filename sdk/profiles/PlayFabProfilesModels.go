package profiles

import "time"
                    
// EffectType 
type EffectType string
  
const (
     EffectTypeAllow EffectType = "Allow"
     EffectTypeDeny EffectType = "Deny"
      )
// EntityDataObject an entity object and its associated meta data.
type EntityDataObjectModel struct {
    // DataObject un-escaped JSON object, if DataAsObject is true.
    DataObject interface{} `json:"DataObject,omitempty"`
    // EscapedDataObject escaped string JSON body of the object, if DataAsObject is default or false.
    EscapedDataObject string `json:"EscapedDataObject,omitempty"`
    // ObjectName name of this object.
    ObjectName string `json:"ObjectName,omitempty"`
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// EntityLineage 
type EntityLineageModel struct {
    // CharacterId the Character Id of the associated entity.
    CharacterId string `json:"CharacterId,omitempty"`
    // GroupId the Group Id of the associated entity.
    GroupId string `json:"GroupId,omitempty"`
    // MasterPlayerAccountId the Master Player Account Id of the associated entity.
    MasterPlayerAccountId string `json:"MasterPlayerAccountId,omitempty"`
    // NamespaceId the Namespace Id of the associated entity.
    NamespaceId string `json:"NamespaceId,omitempty"`
    // TitleId the Title Id of the associated entity.
    TitleId string `json:"TitleId,omitempty"`
    // TitlePlayerAccountId the Title Player Account Id of the associated entity.
    TitlePlayerAccountId string `json:"TitlePlayerAccountId,omitempty"`
}

// EntityPermissionStatement 
type EntityPermissionStatementModel struct {
    // Action the action this statement effects. May be 'Read', 'Write' or '*' for both read and write.
    Action string `json:"Action,omitempty"`
    // Comment a comment about the statement. Intended solely for bookkeeping and debugging.
    Comment string `json:"Comment,omitempty"`
    // Condition additional conditions to be applied for entity resources.
    Condition interface{} `json:"Condition,omitempty"`
    // Effect the effect this statement will have. It may be either Allow or Deny
    Effect EffectType `json:"Effect,omitempty"`
    // Principal the principal this statement will effect.
    Principal interface{} `json:"Principal,omitempty"`
    // Resource the resource this statements effects. Similar to 'pfrn:data--title![Title ID]/Profile/*'
    Resource string `json:"Resource,omitempty"`
}

// EntityProfileBody 
type EntityProfileBodyModel struct {
    // AvatarUrl avatar URL for the entity.
    AvatarUrl string `json:"AvatarUrl,omitempty"`
    // Created the creation time of this profile in UTC.
    Created time.Time `json:"Created,omitempty"`
    // DisplayName the display name of the entity. This field may serve different purposes for different entity types. i.e.: for a title
// player account it could represent the display name of the player, whereas on a character it could be character's name.
    DisplayName string `json:"DisplayName,omitempty"`
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityChain the chain of responsibility for this entity. Use Lineage.
    EntityChain string `json:"EntityChain,omitempty"`
    // Files the files on this profile.
    Files map[string]EntityProfileFileMetadataModel `json:"Files,omitempty"`
    // Language the language on this profile.
    Language string `json:"Language,omitempty"`
    // LeaderboardMetadata leaderboard metadata for the entity.
    LeaderboardMetadata string `json:"LeaderboardMetadata,omitempty"`
    // Lineage the lineage of this profile.
    Lineage *EntityLineageModel `json:"Lineage,omitempty"`
    // Objects the objects on this profile.
    Objects map[string]EntityDataObjectModel `json:"Objects,omitempty"`
    // Permissions the permissions that govern access to this entity profile and its properties. Only includes permissions set on this
// profile, not global statements from titles and namespaces.
    Permissions []EntityPermissionStatementModel `json:"Permissions,omitempty"`
    // Statistics the statistics on this profile.
    Statistics map[string]EntityStatisticValueModel `json:"Statistics,omitempty"`
    // VersionNumber the version number of the profile in persistent storage at the time of the read. Used for optional optimistic
// concurrency during update.
    VersionNumber int32 `json:"VersionNumber,omitempty"`
}

// EntityProfileFileMetadata an entity file's meta data. To get a download URL call File/GetFiles API.
type EntityProfileFileMetadataModel struct {
    // Checksum checksum value for the file
    Checksum string `json:"Checksum,omitempty"`
    // FileName name of the file
    FileName string `json:"FileName,omitempty"`
    // LastModified last UTC time the file was modified
    LastModified time.Time `json:"LastModified,omitempty"`
    // Size storage service's reported byte count
    Size int32 `json:"Size,omitempty"`
}

// EntityStatisticChildValue 
type EntityStatisticChildValueModel struct {
    // ChildName child name value, if child statistic
    ChildName string `json:"ChildName,omitempty"`
    // Metadata child statistic metadata
    Metadata string `json:"Metadata,omitempty"`
    // Value child statistic value
    Value int32 `json:"Value,omitempty"`
}

// EntityStatisticValue 
type EntityStatisticValueModel struct {
    // ChildStatistics child statistic values
    ChildStatistics map[string]EntityStatisticChildValueModel `json:"ChildStatistics,omitempty"`
    // Metadata statistic metadata
    Metadata string `json:"Metadata,omitempty"`
    // Name statistic name
    Name string `json:"Name,omitempty"`
    // Value statistic value
    Value int32 `json:"Value,omitempty"`
    // Version statistic version
    Version int32 `json:"Version,omitempty"`
}

// GetEntityProfileRequest given an entity type and entity identifier will retrieve the profile from the entity store. If the profile being
// retrieved is the caller's, then the read operation is consistent, if not it is an inconsistent read. An inconsistent
// read means that we do not guarantee all committed writes have occurred before reading the profile, allowing for a stale
// read. If consistency is important the Version Number on the result can be used to compare which version of the profile
// any reader has.
type GetEntityProfileRequestModel struct {
    // DataAsObject determines whether the objects will be returned as an escaped JSON string or as a un-escaped JSON object. Default is
// JSON string.
    DataAsObject bool `json:"DataAsObject,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetEntityProfileResponse 
type GetEntityProfileResponseModel struct {
    // Profile entity profile
    Profile *EntityProfileBodyModel `json:"Profile,omitempty"`
}

// GetEntityProfilesRequest given a set of entity types and entity identifiers will retrieve all readable profiles properties for the caller.
// Profiles that the caller is not allowed to read will silently not be included in the results.
type GetEntityProfilesRequestModel struct {
    // DataAsObject determines whether the objects will be returned as an escaped JSON string or as a un-escaped JSON object. Default is
// JSON string.
    DataAsObject bool `json:"DataAsObject,omitempty"`
    // Entities entity keys of the profiles to load. Must be between 1 and 25
    Entities []EntityKeyModel `json:"Entities,omitempty"`
}

// GetEntityProfilesResponse 
type GetEntityProfilesResponseModel struct {
    // Profiles entity profiles
    Profiles []EntityProfileBodyModel `json:"Profiles,omitempty"`
}

// GetGlobalPolicyRequest retrieves the title access policy that is used before the profile's policy is inspected during a request. If never
// customized this will return the default starter policy built by PlayFab.
type GetGlobalPolicyRequestModel struct {
}

// GetGlobalPolicyResponse 
type GetGlobalPolicyResponseModel struct {
    // Permissions the permissions that govern access to all entities under this title or namespace.
    Permissions []EntityPermissionStatementModel `json:"Permissions,omitempty"`
}

// GetTitlePlayersFromMasterPlayerAccountIdsRequest given a master player account id (PlayFab ID), returns all title player accounts associated with it.
type GetTitlePlayersFromMasterPlayerAccountIdsRequestModel struct {
    // MasterPlayerAccountIds master player account ids.
    MasterPlayerAccountIds []string `json:"MasterPlayerAccountIds,omitempty"`
    // TitleId id of title to get players from.
    TitleId string `json:"TitleId,omitempty"`
}

// GetTitlePlayersFromMasterPlayerAccountIdsResponse 
type GetTitlePlayersFromMasterPlayerAccountIdsResponseModel struct {
    // TitlePlayerAccounts dictionary of master player ids mapped to title player entity keys and id pairs
    TitlePlayerAccounts map[string]EntityKeyModel `json:"TitlePlayerAccounts,omitempty"`
}

// OperationTypes 
type OperationTypes string
  
const (
     OperationTypesCreated OperationTypes = "Created"
     OperationTypesUpdated OperationTypes = "Updated"
     OperationTypesDeleted OperationTypes = "Deleted"
     OperationTypesNone OperationTypes = "None"
      )
// SetAvatarUrlRequest given an entity profile, will update its avatar url to the provided url if the profile's version is equal to the
// specified value
type SetAvatarUrlRequestModel struct {
    // AvatarUrl uRL of the avatar image. If empty, it removes the existing avatar URL
    AvatarUrl string `json:"AvatarUrl,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the expected version of the profile to perform this update on
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// SetAvatarUrlResponse 
type SetAvatarUrlResponseModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // OperationResult the type of operation that occurred on the profile's avatar URL
    OperationResult OperationTypes `json:"OperationResult,omitempty"`
    // ProfileVersion the updated version of the profile after the avatar URL update
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
}

// SetDisplayNameRequest given an entity profile, will update its display name to the one passed in if the profile's version is equal to the
// specified value
type SetDisplayNameRequestModel struct {
    // DisplayName the new value to be set on Entity Profile's display name
    DisplayName string `json:"DisplayName,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ExpectedVersion the expected version of a profile to perform this update on
    ExpectedVersion int32 `json:"ExpectedVersion,omitempty"`
}

// SetDisplayNameResponse 
type SetDisplayNameResponseModel struct {
    // OperationResult the type of operation that occured on the profile's display name
    OperationResult OperationTypes `json:"OperationResult,omitempty"`
    // VersionNumber the updated version of the profile after the display name update
    VersionNumber int32 `json:"VersionNumber,omitempty"`
}

// SetEntityProfilePolicyRequest this will set the access policy statements on the given entity profile. This is not additive, any existing statements
// will be replaced with the statements in this request.
type SetEntityProfilePolicyRequestModel struct {
    // Entity the entity to perform this action on.
    Entity EntityKeyModel `json:"Entity,omitempty"`
    // Statements the statements to include in the access policy.
    Statements []EntityPermissionStatementModel `json:"Statements,omitempty"`
}

// SetEntityProfilePolicyResponse 
type SetEntityProfilePolicyResponseModel struct {
    // Permissions the permissions that govern access to this entity profile and its properties. Only includes permissions set on this
// profile, not global statements from titles and namespaces.
    Permissions []EntityPermissionStatementModel `json:"Permissions,omitempty"`
}

// SetGlobalPolicyRequest updates the title access policy that is used before the profile's policy is inspected during a request. Policies are
// compiled and cached for several minutes so an update here may not be reflected in behavior for a short time.
type SetGlobalPolicyRequestModel struct {
    // Permissions the permissions that govern access to all entities under this title or namespace.
    Permissions []EntityPermissionStatementModel `json:"Permissions,omitempty"`
}

// SetGlobalPolicyResponse 
type SetGlobalPolicyResponseModel struct {
}

// SetProfileLanguageRequest given an entity profile, will update its language to the one passed in if the profile's version is equal to the one
// passed in.
type SetProfileLanguageRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ExpectedVersion the expected version of a profile to perform this update on
    ExpectedVersion int32 `json:"ExpectedVersion,omitempty"`
    // Language the language to set on the given entity. Deletes the profile's language if passed in a null string.
    Language string `json:"Language,omitempty"`
}

// SetProfileLanguageResponse 
type SetProfileLanguageResponseModel struct {
    // OperationResult the type of operation that occured on the profile's language
    OperationResult OperationTypes `json:"OperationResult,omitempty"`
    // VersionNumber the updated version of the profile after the language update
    VersionNumber int32 `json:"VersionNumber,omitempty"`
}
