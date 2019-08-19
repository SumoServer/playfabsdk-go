package events

// This code was generated by a tool. Any changes may be overwritten

import (
    "encoding/json"

    playfab "github.com/dgkanatsios/playfabsdk-go/sdk"

    "github.com/mitchellh/mapstructure"
)

// WriteEvents write batches of entity based events to PlayStream.
// https://api.playfab.com/Documentation/Events/method/WriteEvents
func WriteEvents(settings *playfab.Settings, postData *WriteEventsRequestModel, entityToken string) (*WriteEventsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Event/WriteEvents", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &WriteEventsResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}

// WriteTelemetryEvents write batches of entity based events to as Telemetry events (bypass PlayStream).
// https://api.playfab.com/Documentation/Events/method/WriteTelemetryEvents
func WriteTelemetryEvents(settings *playfab.Settings, postData *WriteEventsRequestModel, entityToken string) (*WriteEventsResponseModel, error) {
    if entityToken == "" {
        return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
    }
    b, errMarshal := json.Marshal(postData)
    if errMarshal != nil {
        return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
    }

    sourceMap, err := playfab.Request(settings, b, "/Event/WriteTelemetryEvents", "X-EntityToken", entityToken)
    if err != nil {
        return nil, err
    }
    
    result := &WriteEventsResponseModel{}

    config := mapstructure.DecoderConfig{
        DecodeHook: playfab.StringToDateTimeHook,
        Result:     result,
    }
    
    decoder, errDecoding := mapstructure.NewDecoder(&config)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }
   
    errDecoding = decoder.Decode(sourceMap)
    if errDecoding != nil {
        return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
    }

    return result, nil
}



