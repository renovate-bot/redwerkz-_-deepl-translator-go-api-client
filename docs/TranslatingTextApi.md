# \TranslatingTextApi

All URIs are relative to *https://api.deepl.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TranslateText**](TranslatingTextApi.md#TranslateText) | **Post** /translate | Request Translation



## TranslateText

> TranslateText200Response TranslateText(ctx).Text(text).TargetLang(targetLang).SourceLang(sourceLang).SplitSentences(splitSentences).PreserveFormatting(preserveFormatting).Formality(formality).GlossaryId(glossaryId).TagHandling(tagHandling).NonSplittingTags(nonSplittingTags).OutlineDetection(outlineDetection).SplittingTags(splittingTags).IgnoreTags(ignoreTags).Execute()

Request Translation



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    text := []string{"Inner_example"} // []string | Text to be translated. Only UTF-8-encoded plain text is supported. The parameter may be specified multiple times and translations are returned in the same order as they are requested. Each of the parameter values may contain multiple sentences. Up to 50 texts can be sent for translation in one request.
    targetLang := TODO // TargetLanguage | 
    sourceLang := openapiclient.SourceLanguage("BG") // SourceLanguage |  (optional)
    splitSentences := "splitSentences_example" // string | Sets whether the translation engine should first split the input into sentences. This is enabled by default. Possible values are:  * `0` - no splitting at all, whole input is treated as one sentence  * `1` - splits on punctuation and on newlines  * `nonewlines` - splits on punctuation only, ignoring newlines  For applications that send one sentence per text parameter, it is advisable to set `split_sentences` to `0`, in order to prevent the engine from splitting the sentence unintentionally.   Please note that newlines will split sentences. You should therefore clean files to avoid breaking sentences or set the parameter `split_sentences` to `nonewlines`. In the example below, the two parts of the sentence have been translated separately and this has caused an error: The word \\\"the\\\" has been incorrectly translated as \\\"die\\\" (the feminine definite article in German), though the German word for \\\"sentence\\\", \\\"Satz\\\", is masculine (der Satz).  * Example request: ``` <div>This is the first sentence.</div> ```  * Example response: ``` <div>Dies ist die erste Satz.</div> ``` (optional) (default to "1")
    preserveFormatting := "preserveFormatting_example" // string | Sets whether the translation engine should respect the original formatting, even if it would usually correct some aspects. Possible values are:  * `0`  * `1`  The formatting aspects affected by this setting include:  * Punctuation at the beginning and end of the sentence  * Upper/lower case at the beginning of the sentence (optional) (default to "0")
    formality := TODO // Formality |  (optional) (default to "default")
    glossaryId := "glossaryId_example" // string | Specify the glossary to use for the translation. **Important:** This requires the `source_lang` parameter to be set and the language pair of the glossary has to match the language pair of the request. (optional)
    tagHandling := "tagHandling_example" // string | Sets which kind of tags should be handled. Options currently available:  * `xml`: Enable XML tag handling; see [Handling XML](https://www.deepl.com/docs-api/xml).  * `html`: Enable HTML tag handling; see [Handling HTML (Beta)](https://www.deepl.com/docs-api/html). (optional)
    nonSplittingTags := "nonSplittingTags_example" // string | Comma-separated list of XML tags which never split sentences.   For some XML files, finding tags with textual content and splitting sentences using those tags won't yield the best results. The following example shows the engine splitting sentences on `par` tags and proceeding to translate the parts separately, resulting in an incorrect translation:  * Example request: ``` <par>The firm said it had been </par><par> conducting an internal investigation.</par> ```  * Example response: ``` <par>Die Firma sagte, es sei eine gute Idee gewesen.</par><par> Durchführung einer internen Untersuchung.</par> ```   As this can lead to bad translations, this type of structure should either be avoided, or the `non_splitting_tags` parameter should be set. The following example shows the same call, with the parameter set to `par`:  * Example request: ``` <par>The firm said it had been </par><par> conducting an internal investigation.</par> ```  * Example response: ``` <par>Die Firma sagte, dass sie</par><par> eine interne Untersuchung durchgeführt</par><par> habe</par><par>.</par> ```   This time, the sentence is translated as a whole. The XML tags are now considered markup and copied into the translated sentence. As the translation of the words \\\"had been\\\" has moved to another position in the German sentence, the two par tags are duplicated (which is expected here). (optional)
    outlineDetection := "outlineDetection_example" // string | The automatic detection of the XML structure won't yield best results in all XML files. You can disable this automatic mechanism altogether by setting the `outline_detection` parameter to `0` and selecting the tags that should be considered structure tags. This will split sentences using the `splitting_tags` parameter.   In the example below, we achieve the same results as the automatic engine by disabling automatic detection with `outline_detection=0` and setting the parameters manually to `tag_handling=xml`, `split_sentences=nonewlines`,  and `splitting_tags=par,title`.  * Example request:       <document>       <meta>         <title>A document's title</title>       </meta>       <content>         <par>This is the first sentence. Followed by a second one.</par>         <par>This is the third sentence.</par>       </content>     </document>   * Example response:       <document>       <meta>         <title>Der Titel eines Dokuments</title>       </meta>       <content>         <par>Das ist der erste Satz. Gefolgt von einem zweiten.</par>         <par>Dies ist der dritte Satz.</par>       </content>     </document>   While this approach is slightly more complicated, it allows for greater control over the structure of the translation output. (optional)
    splittingTags := "splittingTags_example" // string | Comma-separated list of XML tags which always cause splits.   See the example in the `outline_detection` parameter's description. (optional)
    ignoreTags := "ignoreTags_example" // string | Comma-separated list of XML tags that indicate text not to be translated.   Use this paramter to ensure that elements in the original text are not altered in the translation (e.g., trademarks, product names) and insert tags into your original text. In the following example, the `ignore_tags` parameter is set to `keep`:  * Example request: ``` Please open the page <keep>Settings</keep> to configure your system. ```  * Example response: ``` Bitte öffnen Sie die Seite <keep>Settings</keep> um Ihr System zu konfigurieren. ``` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TranslatingTextApi.TranslateText(context.Background()).Text(text).TargetLang(targetLang).SourceLang(sourceLang).SplitSentences(splitSentences).PreserveFormatting(preserveFormatting).Formality(formality).GlossaryId(glossaryId).TagHandling(tagHandling).NonSplittingTags(nonSplittingTags).OutlineDetection(outlineDetection).SplittingTags(splittingTags).IgnoreTags(ignoreTags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TranslatingTextApi.TranslateText``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TranslateText`: TranslateText200Response
    fmt.Fprintf(os.Stdout, "Response from `TranslatingTextApi.TranslateText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTranslateTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **[]string** | Text to be translated. Only UTF-8-encoded plain text is supported. The parameter may be specified multiple times and translations are returned in the same order as they are requested. Each of the parameter values may contain multiple sentences. Up to 50 texts can be sent for translation in one request. | 
 **targetLang** | [**TargetLanguage**](TargetLanguage.md) |  | 
 **sourceLang** | [**SourceLanguage**](SourceLanguage.md) |  | 
 **splitSentences** | **string** | Sets whether the translation engine should first split the input into sentences. This is enabled by default. Possible values are:  * &#x60;0&#x60; - no splitting at all, whole input is treated as one sentence  * &#x60;1&#x60; - splits on punctuation and on newlines  * &#x60;nonewlines&#x60; - splits on punctuation only, ignoring newlines  For applications that send one sentence per text parameter, it is advisable to set &#x60;split_sentences&#x60; to &#x60;0&#x60;, in order to prevent the engine from splitting the sentence unintentionally.   Please note that newlines will split sentences. You should therefore clean files to avoid breaking sentences or set the parameter &#x60;split_sentences&#x60; to &#x60;nonewlines&#x60;. In the example below, the two parts of the sentence have been translated separately and this has caused an error: The word \\\&quot;the\\\&quot; has been incorrectly translated as \\\&quot;die\\\&quot; (the feminine definite article in German), though the German word for \\\&quot;sentence\\\&quot;, \\\&quot;Satz\\\&quot;, is masculine (der Satz).  * Example request: &#x60;&#x60;&#x60; &lt;div&gt;This is the first sentence.&lt;/div&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;div&gt;Dies ist die erste Satz.&lt;/div&gt; &#x60;&#x60;&#x60; | [default to &quot;1&quot;]
 **preserveFormatting** | **string** | Sets whether the translation engine should respect the original formatting, even if it would usually correct some aspects. Possible values are:  * &#x60;0&#x60;  * &#x60;1&#x60;  The formatting aspects affected by this setting include:  * Punctuation at the beginning and end of the sentence  * Upper/lower case at the beginning of the sentence | [default to &quot;0&quot;]
 **formality** | [**Formality**](Formality.md) |  | [default to &quot;default&quot;]
 **glossaryId** | **string** | Specify the glossary to use for the translation. **Important:** This requires the &#x60;source_lang&#x60; parameter to be set and the language pair of the glossary has to match the language pair of the request. | 
 **tagHandling** | **string** | Sets which kind of tags should be handled. Options currently available:  * &#x60;xml&#x60;: Enable XML tag handling; see [Handling XML](https://www.deepl.com/docs-api/xml).  * &#x60;html&#x60;: Enable HTML tag handling; see [Handling HTML (Beta)](https://www.deepl.com/docs-api/html). | 
 **nonSplittingTags** | **string** | Comma-separated list of XML tags which never split sentences.   For some XML files, finding tags with textual content and splitting sentences using those tags won&#39;t yield the best results. The following example shows the engine splitting sentences on &#x60;par&#x60; tags and proceeding to translate the parts separately, resulting in an incorrect translation:  * Example request: &#x60;&#x60;&#x60; &lt;par&gt;The firm said it had been &lt;/par&gt;&lt;par&gt; conducting an internal investigation.&lt;/par&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;par&gt;Die Firma sagte, es sei eine gute Idee gewesen.&lt;/par&gt;&lt;par&gt; Durchführung einer internen Untersuchung.&lt;/par&gt; &#x60;&#x60;&#x60;   As this can lead to bad translations, this type of structure should either be avoided, or the &#x60;non_splitting_tags&#x60; parameter should be set. The following example shows the same call, with the parameter set to &#x60;par&#x60;:  * Example request: &#x60;&#x60;&#x60; &lt;par&gt;The firm said it had been &lt;/par&gt;&lt;par&gt; conducting an internal investigation.&lt;/par&gt; &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; &lt;par&gt;Die Firma sagte, dass sie&lt;/par&gt;&lt;par&gt; eine interne Untersuchung durchgeführt&lt;/par&gt;&lt;par&gt; habe&lt;/par&gt;&lt;par&gt;.&lt;/par&gt; &#x60;&#x60;&#x60;   This time, the sentence is translated as a whole. The XML tags are now considered markup and copied into the translated sentence. As the translation of the words \\\&quot;had been\\\&quot; has moved to another position in the German sentence, the two par tags are duplicated (which is expected here). | 
 **outlineDetection** | **string** | The automatic detection of the XML structure won&#39;t yield best results in all XML files. You can disable this automatic mechanism altogether by setting the &#x60;outline_detection&#x60; parameter to &#x60;0&#x60; and selecting the tags that should be considered structure tags. This will split sentences using the &#x60;splitting_tags&#x60; parameter.   In the example below, we achieve the same results as the automatic engine by disabling automatic detection with &#x60;outline_detection&#x3D;0&#x60; and setting the parameters manually to &#x60;tag_handling&#x3D;xml&#x60;, &#x60;split_sentences&#x3D;nonewlines&#x60;,  and &#x60;splitting_tags&#x3D;par,title&#x60;.  * Example request:       &lt;document&gt;       &lt;meta&gt;         &lt;title&gt;A document&#39;s title&lt;/title&gt;       &lt;/meta&gt;       &lt;content&gt;         &lt;par&gt;This is the first sentence. Followed by a second one.&lt;/par&gt;         &lt;par&gt;This is the third sentence.&lt;/par&gt;       &lt;/content&gt;     &lt;/document&gt;   * Example response:       &lt;document&gt;       &lt;meta&gt;         &lt;title&gt;Der Titel eines Dokuments&lt;/title&gt;       &lt;/meta&gt;       &lt;content&gt;         &lt;par&gt;Das ist der erste Satz. Gefolgt von einem zweiten.&lt;/par&gt;         &lt;par&gt;Dies ist der dritte Satz.&lt;/par&gt;       &lt;/content&gt;     &lt;/document&gt;   While this approach is slightly more complicated, it allows for greater control over the structure of the translation output. | 
 **splittingTags** | **string** | Comma-separated list of XML tags which always cause splits.   See the example in the &#x60;outline_detection&#x60; parameter&#39;s description. | 
 **ignoreTags** | **string** | Comma-separated list of XML tags that indicate text not to be translated.   Use this paramter to ensure that elements in the original text are not altered in the translation (e.g., trademarks, product names) and insert tags into your original text. In the following example, the &#x60;ignore_tags&#x60; parameter is set to &#x60;keep&#x60;:  * Example request: &#x60;&#x60;&#x60; Please open the page &lt;keep&gt;Settings&lt;/keep&gt; to configure your system. &#x60;&#x60;&#x60;  * Example response: &#x60;&#x60;&#x60; Bitte öffnen Sie die Seite &lt;keep&gt;Settings&lt;/keep&gt; um Ihr System zu konfigurieren. &#x60;&#x60;&#x60; | 

### Return type

[**TranslateText200Response**](TranslateText200Response.md)

### Authorization

[auth_header](../README.md#auth_header)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

