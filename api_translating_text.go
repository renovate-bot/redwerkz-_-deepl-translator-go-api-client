/*
DeepL API

The DeepL API provides programmatic access to DeepL’s machine translation technology.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)


// TranslatingTextApiService TranslatingTextApi service
type TranslatingTextApiService service

type ApiTranslateTextRequest struct {
	ctx context.Context
	ApiService *TranslatingTextApiService
	text *[]string
	targetLang *TargetLanguage
	sourceLang *SourceLanguage
	splitSentences *string
	preserveFormatting *string
	formality *Formality
	glossaryId *string
	tagHandling *string
	nonSplittingTags *string
	outlineDetection *string
	splittingTags *string
	ignoreTags *string
}

// Text to be translated. Only UTF-8-encoded plain text is supported. The parameter may be specified multiple times and translations are returned in the same order as they are requested. Each of the parameter values may contain multiple sentences. Up to 50 texts can be sent for translation in one request.
func (r ApiTranslateTextRequest) Text(text []string) ApiTranslateTextRequest {
	r.text = &text
	return r
}

func (r ApiTranslateTextRequest) TargetLang(targetLang TargetLanguage) ApiTranslateTextRequest {
	r.targetLang = &targetLang
	return r
}

func (r ApiTranslateTextRequest) SourceLang(sourceLang SourceLanguage) ApiTranslateTextRequest {
	r.sourceLang = &sourceLang
	return r
}

// Sets whether the translation engine should first split the input into sentences. This is enabled by default. Possible values are:  * &#x60;0&#x60; - no splitting at all, whole input is treated as one sentence  * &#x60;1&#x60; - splits on punctuation and on newlines  * &#x60;nonewlines&#x60; - splits on punctuation only, ignoring newlines  For applications that send one sentence per text parameter, it is advisable to set &#x60;split_sentences&#x60; to &#x60;0&#x60;, in order to prevent the engine from splitting the sentence unintentionally.   Please note that newlines will split sentences. You should therefore clean files to avoid breaking sentences or set the parameter &#x60;split_sentences&#x60; to &#x60;nonewlines&#x60;. In the example below, the two parts of the sentence have been translated separately and this has caused an error: The word \\\&quot;the\\\&quot; has been incorrectly translated as \\\&quot;die\\\&quot; (the feminine definite article in German), though the German word for \\\&quot;sentence\\\&quot;, \\\&quot;Satz\\\&quot;, is masculine (der Satz).  * Example request: &#x60;&#x60;&#x60; &lt;div&gt;This is the first sentence.&lt;/div&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;div&gt;Dies ist die erste Satz.&lt;/div&gt; &#x60;&#x60;&#x60;
func (r ApiTranslateTextRequest) SplitSentences(splitSentences string) ApiTranslateTextRequest {
	r.splitSentences = &splitSentences
	return r
}

// Sets whether the translation engine should respect the original formatting, even if it would usually correct some aspects. Possible values are:  * &#x60;0&#x60;  * &#x60;1&#x60;  The formatting aspects affected by this setting include:  * Punctuation at the beginning and end of the sentence  * Upper/lower case at the beginning of the sentence
func (r ApiTranslateTextRequest) PreserveFormatting(preserveFormatting string) ApiTranslateTextRequest {
	r.preserveFormatting = &preserveFormatting
	return r
}

func (r ApiTranslateTextRequest) Formality(formality Formality) ApiTranslateTextRequest {
	r.formality = &formality
	return r
}

// Specify the glossary to use for the translation. **Important:** This requires the &#x60;source_lang&#x60; parameter to be set and the language pair of the glossary has to match the language pair of the request.
func (r ApiTranslateTextRequest) GlossaryId(glossaryId string) ApiTranslateTextRequest {
	r.glossaryId = &glossaryId
	return r
}

// Sets which kind of tags should be handled. Options currently available:  * &#x60;xml&#x60;: Enable XML tag handling; see [Handling XML](https://www.deepl.com/docs-api/xml).  * &#x60;html&#x60;: Enable HTML tag handling; see [Handling HTML (Beta)](https://www.deepl.com/docs-api/html).
func (r ApiTranslateTextRequest) TagHandling(tagHandling string) ApiTranslateTextRequest {
	r.tagHandling = &tagHandling
	return r
}

// Comma-separated list of XML tags which never split sentences.   For some XML files, finding tags with textual content and splitting sentences using those tags won&#39;t yield the best results. The following example shows the engine splitting sentences on &#x60;par&#x60; tags and proceeding to translate the parts separately, resulting in an incorrect translation:  * Example request: &#x60;&#x60;&#x60; &lt;par&gt;The firm said it had been &lt;/par&gt;&lt;par&gt; conducting an internal investigation.&lt;/par&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;par&gt;Die Firma sagte, es sei eine gute Idee gewesen.&lt;/par&gt;&lt;par&gt; Durchführung einer internen Untersuchung.&lt;/par&gt; &#x60;&#x60;&#x60;   As this can lead to bad translations, this type of structure should either be avoided, or the &#x60;non_splitting_tags&#x60; parameter should be set. The following example shows the same call, with the parameter set to &#x60;par&#x60;:  * Example request: &#x60;&#x60;&#x60; &lt;par&gt;The firm said it had been &lt;/par&gt;&lt;par&gt; conducting an internal investigation.&lt;/par&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;par&gt;Die Firma sagte, dass sie&lt;/par&gt;&lt;par&gt; eine interne Untersuchung durchgeführt&lt;/par&gt;&lt;par&gt; habe&lt;/par&gt;&lt;par&gt;.&lt;/par&gt; &#x60;&#x60;&#x60;   This time, the sentence is translated as a whole. The XML tags are now considered markup and copied into the translated sentence. As the translation of the words \\\&quot;had been\\\&quot; has moved to another position in the German sentence, the two par tags are duplicated (which is expected here).
func (r ApiTranslateTextRequest) NonSplittingTags(nonSplittingTags string) ApiTranslateTextRequest {
	r.nonSplittingTags = &nonSplittingTags
	return r
}

// The automatic detection of the XML structure won&#39;t yield best results in all XML files. You can disable this automatic mechanism altogether by setting the &#x60;outline_detection&#x60; parameter to &#x60;0&#x60; and selecting the tags that should be considered structure tags. This will split sentences using the &#x60;splitting_tags&#x60; parameter.   In the example below, we achieve the same results as the automatic engine by disabling automatic detection with &#x60;outline_detection&#x3D;0&#x60; and setting the parameters manually to &#x60;tag_handling&#x3D;xml&#x60;, &#x60;split_sentences&#x3D;nonewlines&#x60;,  and &#x60;splitting_tags&#x3D;par,title&#x60;.  * Example request:       &lt;document&gt;       &lt;meta&gt;         &lt;title&gt;A document&#39;s title&lt;/title&gt;       &lt;/meta&gt;       &lt;content&gt;         &lt;par&gt;This is the first sentence. Followed by a second one.&lt;/par&gt;         &lt;par&gt;This is the third sentence.&lt;/par&gt;       &lt;/content&gt;     &lt;/document&gt;   * Example response:       &lt;document&gt;       &lt;meta&gt;         &lt;title&gt;Der Titel eines Dokuments&lt;/title&gt;       &lt;/meta&gt;       &lt;content&gt;         &lt;par&gt;Das ist der erste Satz. Gefolgt von einem zweiten.&lt;/par&gt;         &lt;par&gt;Dies ist der dritte Satz.&lt;/par&gt;       &lt;/content&gt;     &lt;/document&gt;   While this approach is slightly more complicated, it allows for greater control over the structure of the translation output.
func (r ApiTranslateTextRequest) OutlineDetection(outlineDetection string) ApiTranslateTextRequest {
	r.outlineDetection = &outlineDetection
	return r
}

// Comma-separated list of XML tags which always cause splits.   See the example in the &#x60;outline_detection&#x60; parameter&#39;s description.
func (r ApiTranslateTextRequest) SplittingTags(splittingTags string) ApiTranslateTextRequest {
	r.splittingTags = &splittingTags
	return r
}

// Comma-separated list of XML tags that indicate text not to be translated.   Use this paramter to ensure that elements in the original text are not altered in the translation (e.g., trademarks, product names) and insert tags into your original text. In the following example, the &#x60;ignore_tags&#x60; parameter is set to &#x60;keep&#x60;:  * Example request: &#x60;&#x60;&#x60; Please open the page &lt;keep&gt;Settings&lt;/keep&gt; to configure your system. &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; Bitte öffnen Sie die Seite &lt;keep&gt;Settings&lt;/keep&gt; um Ihr System zu konfigurieren. &#x60;&#x60;&#x60;
func (r ApiTranslateTextRequest) IgnoreTags(ignoreTags string) ApiTranslateTextRequest {
	r.ignoreTags = &ignoreTags
	return r
}

func (r ApiTranslateTextRequest) Execute() (*TranslateText200Response, *http.Response, error) {
	return r.ApiService.TranslateTextExecute(r)
}

/*
TranslateText Request Translation

The translate function.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTranslateTextRequest
*/
func (a *TranslatingTextApiService) TranslateText(ctx context.Context) ApiTranslateTextRequest {
	return ApiTranslateTextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TranslateText200Response
func (a *TranslatingTextApiService) TranslateTextExecute(r ApiTranslateTextRequest) (*TranslateText200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TranslateText200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslatingTextApiService.TranslateText")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/translate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.text == nil {
		return localVarReturnValue, nil, reportError("text is required and must be specified")
	}
	if len(*r.text) > 50 {
		return localVarReturnValue, nil, reportError("text must have less than 50 elements")
	}
	if r.targetLang == nil {
		return localVarReturnValue, nil, reportError("targetLang is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	localVarFormParams.Add("text", parameterToString(*r.text, "multi"))
	if r.sourceLang != nil {
		localVarFormParams.Add("source_lang", parameterToString(*r.sourceLang, ""))
	}
	localVarFormParams.Add("target_lang", parameterToString(*r.targetLang, ""))
	if r.splitSentences != nil {
		localVarFormParams.Add("split_sentences", parameterToString(*r.splitSentences, ""))
	}
	if r.preserveFormatting != nil {
		localVarFormParams.Add("preserve_formatting", parameterToString(*r.preserveFormatting, ""))
	}
	if r.formality != nil {
		localVarFormParams.Add("formality", parameterToString(*r.formality, ""))
	}
	if r.glossaryId != nil {
		localVarFormParams.Add("glossary_id", parameterToString(*r.glossaryId, ""))
	}
	if r.tagHandling != nil {
		localVarFormParams.Add("tag_handling", parameterToString(*r.tagHandling, ""))
	}
	if r.nonSplittingTags != nil {
		localVarFormParams.Add("non_splitting_tags", parameterToString(*r.nonSplittingTags, ""))
	}
	if r.outlineDetection != nil {
		localVarFormParams.Add("outline_detection", parameterToString(*r.outlineDetection, ""))
	}
	if r.splittingTags != nil {
		localVarFormParams.Add("splitting_tags", parameterToString(*r.splittingTags, ""))
	}
	if r.ignoreTags != nil {
		localVarFormParams.Add("ignore_tags", parameterToString(*r.ignoreTags, ""))
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["auth_header"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
