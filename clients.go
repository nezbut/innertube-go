package innertubego

import (
	"context"
	"errors"
	"net/http"
)

// InnerTube struct
type InnerTube struct {
	Adaptor Adaptor
}

// Adaptor interface
type Adaptor interface {
	Dispatch(ctx context.Context, endpoint string, params map[string]string, body map[string]interface{}) (map[string]interface{}, error)
}

// NewInnerTube creates a new InnerTube instance
func NewInnerTube(httpClient *http.Client, clientName, clientVersion string, apiKey, userAgent, referer string, locale *Locale, auto bool) (*InnerTube, error) {
	if clientName == "" {
		return nil, errors.New("Precondition failed: Missing client name")
	}

	if clientVersion == "" {
		return nil, errors.New("Precondition failed: Missing client version")
	}
	var context ClientContext
	if auto {
		context = GetContext(clientName)
	} else {
		context = ClientContext{
			ClientName:    clientName,
			ClientVersion: clientVersion,
			APIKey:        apiKey,
			UserAgent:     userAgent,
			Referer:       referer,
			Locale:        locale,
		}
	}

	return &InnerTube{
		Adaptor: NewInnerTubeAdaptor(context, httpClient),
	}, nil
}

// Call method to make requests
func (it *InnerTube) Call(ctx context.Context, endpoint string, params map[string]string, body map[string]interface{}) (map[string]interface{}, error) {
	response, err := it.Adaptor.Dispatch(ctx, endpoint, params, body)
	if err != nil {
		return nil, err
	}

	delete(response, "responseContext") // Remove responseContext if exists
	return response, nil
}

// Example API call methods
func (it *InnerTube) Config(ctx context.Context) (map[string]interface{}, error) {
	return it.Call(ctx, "CONFIG", nil, nil)
}

func (it *InnerTube) Guide(ctx context.Context) (map[string]interface{}, error) {
	return it.Call(ctx, "GUIDE", nil, nil)
}

func (it *InnerTube) Player(ctx context.Context, videoID string) (map[string]interface{}, error) {
	return it.Call(ctx, "PLAYER", nil, Filter(map[string]interface{}{
		"videoId": videoID,
	}))
}

func (it *InnerTube) Browse(ctx context.Context, browseID *string, params *string, continuation *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"browseId":     browseID,
		"params":       params,
		"continuation": continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "BROWSE", nil, Filter(body))
}

func (it *InnerTube) Search(ctx context.Context, query *string, params *string, continuation *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"query":        query,
		"params":       params,
		"continuation": continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "SEARCH", nil, Filter(body))
}

func (it *InnerTube) Next(ctx context.Context, videoId *string, playlistId *string, params *string, index *int, continuation *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"videoId":       videoId,
		"playlistId":    playlistId,
		"params":        params,
		"playlistIndex": index,
		"continuation":  continuation,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "NEXT", nil, Filter(body))
}

func (it *InnerTube) GetTranscript(ctx context.Context, params *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"params": params,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "GET_TRANSCRIPT", nil, Filter(body))
}

func (it *InnerTube) MusicGetSearchSuggestions(ctx context.Context, input *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"input": input,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "MUSIC/GET_SEARCH_SUGGESTIONS", nil, Filter(body))
}

func (it *InnerTube) MusicGetQueue(ctx context.Context, videoIds *[]string, playlistId *string) (map[string]interface{}, error) {
	body := map[string]interface{}{
		"playlistId": playlistId,
		"videoIds":   videoIds,
	}
	//fmt.Println("body: ", body)
	//fmt.Println("Filter(body): ", Filter(body))
	return it.Call(ctx, "MUSIC/GET_QUEUE", nil, Filter(body))
}
