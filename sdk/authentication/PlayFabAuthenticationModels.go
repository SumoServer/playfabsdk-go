package authentication

import "time"
                    
// ActivateAPIKeyRequest 
type ActivateAPIKeyRequestModel struct {
    // APIKeyId unique identifier for the entity API key to activate.
    APIKeyId string `json:"APIKeyId,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// ActivateAPIKeyResponse 
type ActivateAPIKeyResponseModel struct {
}

// CreateAPIKeyDetails 
type CreateAPIKeyDetailsModel struct {
    // Active whether the key is active for authentication. Inactive keys cannot be used to authenticate.Keys can be activated or
// deactivate using the ActivateKey and DeactivateKey APIs.Deactivating a key is a way to verify that the key is not in use
// before deleting it.
    Active bool `json:"Active,omitempty"`
    // APIKeyId unique identifier for the entity API key. Set in the "X - EntityAPIKeyId" in authentication requests.
    APIKeyId string `json:"APIKeyId,omitempty"`
    // APIKeySecret secret used to authenticate requests with the key. Set in the "X - EntityAPIKeyId" in authentication requests.The secret
// value is returned only once in this response and cannot be retrieved afterward, either via API or in Game Manager.API
// key secrets should be stored securely only on trusted servers and never distributed in code or configuration to
// untrusted clients.
    APIKeySecret string `json:"APIKeySecret,omitempty"`
    // Created the time the API key was generated, in UTC.
    Created time.Time `json:"Created,omitempty"`
}

// CreateAPIKeyRequest 
type CreateAPIKeyRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// CreateAPIKeyResponse 
type CreateAPIKeyResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Key the created API key
    Key *CreateAPIKeyDetailsModel `json:"Key,omitempty"`
}

// DeactivateAPIKeyRequest 
type DeactivateAPIKeyRequestModel struct {
    // APIKeyId unique identifier for the entity API key to activate.
    APIKeyId string `json:"APIKeyId,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// DeactivateAPIKeyResponse 
type DeactivateAPIKeyResponseModel struct {
}

// DeleteAPIKeyRequest 
type DeleteAPIKeyRequestModel struct {
    // APIKeyId unique identifier for the entity API key to delete.
    APIKeyId string `json:"APIKeyId,omitempty"`
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// DeleteAPIKeyResponse 
type DeleteAPIKeyResponseModel struct {
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

// GetAPIKeyDetails 
type GetAPIKeyDetailsModel struct {
    // Active whether the key is active for authentication. Inactive keys cannot be used to authenticate.Keys can be activated or
// deactivate using the SetAPIActivation API.Deactivating a key is a way to verify that the key is not in use be before
// deleting it.
    Active bool `json:"Active,omitempty"`
    // APIKeyId unique identifier for the entity API key. Set in the "X - EntityAPIKeyId" in authentication requests.
    APIKeyId string `json:"APIKeyId,omitempty"`
    // Created the time the API key was generated, in UTC.
    Created time.Time `json:"Created,omitempty"`
}

// GetAPIKeysRequest 
type GetAPIKeysRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetAPIKeysResponse 
type GetAPIKeysResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Keys the API keys associated with the given entity.
    Keys []GetAPIKeyDetailsModel `json:"Keys,omitempty"`
}

// GetEntityTokenRequest this API must be called with X-SecretKey, X-Authentication or X-EntityToken headers. An optional EntityKey may be
// included to attempt to set the resulting EntityToken to a specific entity, however the entity must be a relation of the
// caller, such as the master_player_account of a character. If sending X-EntityToken the account will be marked as freshly
// logged in and will issue a new token. If using X-Authentication or X-EntityToken the header must still be valid and
// cannot be expired or revoked.
type GetEntityTokenRequestModel struct {
    // Entity the entity to perform this action on.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
}

// GetEntityTokenResponse 
type GetEntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // EntityToken the token used to set X-EntityToken for all entity based API calls.
    EntityToken string `json:"EntityToken,omitempty"`
    // TokenExpiration the time the token will expire, if it is an expiring token, in UTC.
    TokenExpiration time.Time `json:"TokenExpiration,omitempty"`
}

// ValidateEntityTokenRequest given an entity token, validates that it hasn't exipired or been revoked and will return details of the owner.
type ValidateEntityTokenRequestModel struct {
    // EntityToken client EntityToken
    EntityToken string `json:"EntityToken,omitempty"`
}

// ValidateEntityTokenResponse 
type ValidateEntityTokenResponseModel struct {
    // Entity the entity id and type.
    Entity *EntityKeyModel `json:"Entity,omitempty"`
    // Lineage the lineage of this profile.
    Lineage *EntityLineageModel `json:"Lineage,omitempty"`
}
