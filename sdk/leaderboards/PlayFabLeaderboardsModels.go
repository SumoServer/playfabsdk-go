package leaderboards

import "time"
                    
// CreateStatisticDefinitionRequest 
type CreateStatisticDefinitionRequestModel struct {
    // AggregationMethod aggregation method for this statistic. Default is Last.
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // LeaderboardDefinition additional configuration options for statistics to be used for leaderboards.
    LeaderboardDefinition *LeaderboardDefinitionModel `json:"LeaderboardDefinition,omitempty"`
    // Name immutable name of the statistic. Must be less than 50 characters. Restricted to a-Z, 0-9, '(', ')', '_', '-' and '.'.
    Name string `json:"Name,omitempty"`
}

// DeleteStatisticDefinitionRequest 
type DeleteStatisticDefinitionRequestModel struct {
    // Name name of the statistic to delete.
    Name string `json:"Name,omitempty"`
}

// DeleteStatisticsRequest 
type DeleteStatisticsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ExpectedProfileVersion optional field used for concurrency control. By specifying the previously returned value of ProfileVersion from
// GetProfile API, you can ensure that the object set will only be performed if the profile has not been updated by any
// other clients since the version you last loaded.
    ExpectedProfileVersion int32 `json:"ExpectedProfileVersion,omitempty"`
    // Statistics collection of statistics to remove from this entity.
    Statistics []StatisticDeleteModel `json:"Statistics,omitempty"`
}

// DeleteStatisticsResponse 
type DeleteStatisticsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // Statistics updated entity profile statistics.
    Statistics map[string]EntityStatisticValueModel `json:"Statistics,omitempty"`
}

// EmptyResponse 
type EmptyResponseModel struct {
}

// EntityKey combined entity type and ID structure which uniquely identifies a single entity.
type EntityKeyModel struct {
    // Id unique ID of the entity.
    Id string `json:"Id,omitempty"`
    // Type entity type. See https://api.playfab.com/docs/tutorials/entities/entitytypes
    Type string `json:"Type,omitempty"`
}

// EntityLeaderboardEntry individual rank of an entity in a leaderboard
type EntityLeaderboardEntryModel struct {
    // DisplayName entity's display name.
    DisplayName string `json:"DisplayName,omitempty"`
    // Entity entity identifier.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityLeaderboardMetadata custom entity metadata as set during the last UpdateStatistic API call.
    EntityLeaderboardMetadata string `json:"EntityLeaderboardMetadata,omitempty"`
    // EntityLineage lineage for this entity.
    EntityLineage *EntityLineageModel `json:"EntityLineage,omitempty"`
    // Metadata statistic or statistic child metadata as set during the UpdateStatistic API call.
    Metadata string `json:"Metadata,omitempty"`
    // Rank position on the leaderboard.
    Rank int32 `json:"Rank,omitempty"`
    // Score statistic value.
    Score int32 `json:"Score,omitempty"`
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

// GetEntityLeaderboardRequest request to load a leaderboard.
type GetEntityLeaderboardRequestModel struct {
    // ChildName optional child statistic name to filter the leaderboard to.
    ChildName string `json:"ChildName,omitempty"`
    // EntityType type of entity to get a leaderboard for.
    EntityType string `json:"EntityType,omitempty"`
    // MaxResults maximum number of results to return from the leaderboard. Minimum 1, maximum 1,000.
    MaxResults uint32 `json:"MaxResults,omitempty"`
    // StartingPosition index position to start from. 0 is beginning of leaderboard.
    StartingPosition uint32 `json:"StartingPosition,omitempty"`
    // StatisticName name of the leaderboard statistic.
    StatisticName string `json:"StatisticName,omitempty"`
    // StatisticVersion optional version of the leaderboard statistic, defaults to current version.
    StatisticVersion uint32 `json:"StatisticVersion,omitempty"`
}

// GetEntityLeaderboardResponse leaderboard response
type GetEntityLeaderboardResponseModel struct {
    // Rankings individual entity rankings in the leaderboard, in sorted order by rank.
    Rankings []EntityLeaderboardEntryModel `json:"Rankings,omitempty"`
    // StatisticVersion version of the leaderboard being returned.
    StatisticVersion uint32 `json:"StatisticVersion,omitempty"`
}

// GetLeaderboardAroundEntityRequest request to load a section of a leaderboard centered on a specific entity.
type GetLeaderboardAroundEntityRequestModel struct {
    // ChildName optional statistic child to filter the leaderboard by.
    ChildName string `json:"ChildName,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Offset number of ranks above and below the entity to return. If 5 then at most 11 results will be returned, 5 above, the
// entity, then 5 below. If the entity is at the top then only 6 will be returned, the entity and 5 below.
    Offset uint32 `json:"Offset,omitempty"`
    // StatisticName name of the statistic.
    StatisticName string `json:"StatisticName,omitempty"`
    // StatisticVersion optional version of the statistic, defaults to current.
    StatisticVersion uint32 `json:"StatisticVersion,omitempty"`
}

// GetLeaderboardForEntitiesRequest request a leaderboard limited to a collection of entities.
type GetLeaderboardForEntitiesRequestModel struct {
    // ChildName optional statistic child name to filter the leaderboard by.
    ChildName string `json:"ChildName,omitempty"`
    // Entities collection of Entity IDs to include in the leaderboard.
    Entities []string `json:"Entities,omitempty"`
    // EntityType entity type of all Entity IDs.
    EntityType string `json:"EntityType,omitempty"`
    // StatisticName name of the statistic.
    StatisticName string `json:"StatisticName,omitempty"`
    // StatisticVersion optional version of the statistic, defaults to current.
    StatisticVersion uint32 `json:"StatisticVersion,omitempty"`
}

// GetStatisticDefinitionRequest 
type GetStatisticDefinitionRequestModel struct {
    // Name immutable name of the statistic. Must be less than 50 characters.
    Name string `json:"Name,omitempty"`
}

// GetStatisticDefinitionResponse 
type GetStatisticDefinitionResponseModel struct {
    // AggregationMethod aggregation method of the statistic.
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // Created created time, in UTC
    Created time.Time `json:"Created,omitempty"`
    // LastResetTime last time, in UTC, statistic version was incremented.
    LastResetTime time.Time `json:"LastResetTime,omitempty"`
    // LeaderboardDefinition statistic leaderboard definition.
    LeaderboardDefinition *LeaderboardDefinitionModel `json:"LeaderboardDefinition,omitempty"`
    // Name name of the statistic.
    Name string `json:"Name,omitempty"`
    // Version statistic version.
    Version uint32 `json:"Version,omitempty"`
}

// GetStatisticDefinitionsRequest 
type GetStatisticDefinitionsRequestModel struct {
}

// GetStatisticDefinitionsResponse 
type GetStatisticDefinitionsResponseModel struct {
    // StatisticDefinitions list of statistic definitions for the title.
    StatisticDefinitions []StatisticDefinitionModel `json:"StatisticDefinitions,omitempty"`
}

// IncrementStatisticVersionRequest 
type IncrementStatisticVersionRequestModel struct {
    // Name name of the statistic to increment the version of.
    Name string `json:"Name,omitempty"`
}

// IncrementStatisticVersionResponse 
type IncrementStatisticVersionResponseModel struct {
    // Version new statistic version.
    Version uint32 `json:"Version,omitempty"`
    // VersionToBeDeleted statistic versions and leaderboards to be removed.
    VersionToBeDeleted uint32 `json:"VersionToBeDeleted,omitempty"`
}

// LeaderboardDefinition 
type LeaderboardDefinitionModel struct {
    // ChildLeaderboardNames for non-dynamic child leaderboards the statistic child name must be defined to prevent leaderboard length restrictions.
    ChildLeaderboardNames []string `json:"ChildLeaderboardNames,omitempty"`
    // DynamicChildLeaderboardMaxSize number of entities to rank in a leaderboard when the child is dynamically defined, new entities will not be ranked
// unless that beat an existing score on a leaderboard. Default 0, Maximum 100.
    DynamicChildLeaderboardMaxSize uint32 `json:"DynamicChildLeaderboardMaxSize,omitempty"`
    // ProvisionLeaderboard flag to indicate that leaderboards should be created/updated when this statistic is updated on an entity. Default is
// false.
    ProvisionLeaderboard bool `json:"ProvisionLeaderboard,omitempty"`
    // SortDirection sort direction of the leaderboard, cannot be changed after creation. Default is Descending
    SortDirection LeaderboardSortDirection `json:"SortDirection,omitempty"`
}

// LeaderboardSortDirection 
type LeaderboardSortDirection string
  
const (
     LeaderboardSortDirectionDescending LeaderboardSortDirection = "Descending"
     LeaderboardSortDirectionAscending LeaderboardSortDirection = "Ascending"
      )
// StatisticAggregationMethod 
type StatisticAggregationMethod string
  
const (
     StatisticAggregationMethodLast StatisticAggregationMethod = "Last"
     StatisticAggregationMethodMin StatisticAggregationMethod = "Min"
     StatisticAggregationMethodMax StatisticAggregationMethod = "Max"
     StatisticAggregationMethodSum StatisticAggregationMethod = "Sum"
      )
// StatisticDefinition 
type StatisticDefinitionModel struct {
    // AggregationMethod aggregation method of the statistic.
    AggregationMethod StatisticAggregationMethod `json:"AggregationMethod,omitempty"`
    // Created created time, in UTC
    Created time.Time `json:"Created,omitempty"`
    // LastResetTime last time, in UTC, statistic version was incremented.
    LastResetTime time.Time `json:"LastResetTime,omitempty"`
    // LeaderboardDefinition statistic leaderboard definition.
    LeaderboardDefinition *LeaderboardDefinitionModel `json:"LeaderboardDefinition,omitempty"`
    // Name name of the statistic.
    Name string `json:"Name,omitempty"`
    // Version statistic version.
    Version uint32 `json:"Version,omitempty"`
}

// StatisticDelete 
type StatisticDeleteModel struct {
    // ChildName optionally delete only the child statistic from this entity.
    ChildName string `json:"ChildName,omitempty"`
    // Name name of the statistic, as originally configured.
    Name string `json:"Name,omitempty"`
    // Version optional field to indicate the version of the statistic to set. When empty defaults to the statistic's current version.
    Version uint32 `json:"Version,omitempty"`
}

// StatisticUpdate 
type StatisticUpdateModel struct {
    // ChildName arbitrary statistic child name, using child statistics will cause entity to be ranked on only the child leaderboard. See
// documentation for full details. Must be less than 50 UTF8 encoded characters.
    ChildName string `json:"ChildName,omitempty"`
    // Metadata arbitrary metadata to store along side the statistic or child statistic, will be returned by all Leaderboard APIs. Must
// be less than 50 UTF8 encoded characters.
    Metadata string `json:"Metadata,omitempty"`
    // Name name of the statistic, as originally configured.
    Name string `json:"Name,omitempty"`
    // Value statistic value for the entity. This will be used in accordance with the aggregation method configured for the
// statistics.
    Value int32 `json:"Value,omitempty"`
    // Version optional field to indicate the version of the statistic to set. When empty defaults to the statistic's current version.
    Version uint32 `json:"Version,omitempty"`
}

// UpdateStatisticsRequest 
type UpdateStatisticsRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityLeaderboardMetadata arbitrary metadata to store on the entity that will be returned on all leaderboard APIs. Max length 50 characters.
    EntityLeaderboardMetadata string `json:"EntityLeaderboardMetadata,omitempty"`
    // ExpectedProfileVersion optional field used for concurrency control. By specifying the previously returned value of ProfileVersion from
// GetProfile API, you can ensure that the statistic updates will only be performed if the profile has not been updated by
// any other clients since the version you last loaded.
    ExpectedProfileVersion int32 `json:"ExpectedProfileVersion,omitempty"`
    // ForceUpdate flag to force the values in the request to bypass the aggregation method and be the new value for the statistic. When
// empty defaults to false.
    ForceUpdate bool `json:"ForceUpdate,omitempty"`
    // Statistics collection of statistics to update, maximum 50.
    Statistics []StatisticUpdateModel `json:"Statistics,omitempty"`
}

// UpdateStatisticsResponse 
type UpdateStatisticsResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // ProfileVersion the current version of the profile, can be used for concurrency control during updates.
    ProfileVersion int32 `json:"ProfileVersion,omitempty"`
    // Statistics updated entity profile statistics.
    Statistics map[string]EntityStatisticValueModel `json:"Statistics,omitempty"`
}
