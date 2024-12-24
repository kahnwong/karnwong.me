+++
title = "Elasticsearch with custom dictionary"
date = "2021-05-03"
path = "/posts/2021/05/elasticsearch-with-custom-dictionary"

[taxonomies]
categories = ["software-engineering"]
tags = []

+++

Elasticsearch is a search engine with built-in analyzers (combination of tokenizer and filters), which makes it easier to set it up and get it running, seeing you don’t have to implement NLP logic from scratch. However, for some languages such as Thai, the built-in Thai analyzer may not be working quite as expected.

For instance, for region name search autocomplete, it doesn’t recommend anything when I type `เชียง`, but it should be showing `เชียงใหม่` or `เชียงราย`. This is because these two regions are recognized as one token, which is why it doesn’t recommend anything when querying with `เชียง`.

But if I create a custom dictionary for tokenizers with `เชียง` as one of the tokens, it manages to recommend the two regions when querying with the prefix.

Below is an `index_config` for using a custom dictionary for tokenizer:

```python
{
    "settings": {
        "analysis": {
            "analyzer": {
                "thai_dictionary": {
                    "tokenizer": "standard",
                    "filter": ["char_limit_dictionary"],
                }
            },
            "filter": {
                "char_limit_dictionary": {
                    "type": "dictionary_decompounder",
                    "word_list": tokens,  # <-- word list array here
                    "max_subword_size": 22,
                }
            },
        }
    },
    "mappings": {
        "properties": {
            "title": {  # <-- search key
                "type": "text",
                "analyzer": "thai_dictionary",
            }
        }
    },
}
```

See elasticsearch documentation for more details: <https://www.elastic.co/guide/en/elasticsearch/reference/7.12/analysis-dict-decomp-tokenfilter.html>
