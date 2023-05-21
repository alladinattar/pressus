from natasha import (
    Segmenter,
    MorphVocab,
    NewsEmbedding,
    NewsMorphTagger,
    NewsSyntaxParser,
    NewsNERTagger,
    PER,
    NamesExtractor,
    Doc
)

from flask import Flask, request, jsonify

app = Flask(__name__)


@app.route("/ner", methods=['POST'])
def hello_world():
    body = request.json
    segmenter = Segmenter()
    emb = NewsEmbedding()
    ner_tagger = NewsNERTagger(emb)
    text = body['data']
    doc = Doc(text)
    doc.segment(segmenter)
    doc.tag_ner(ner_tagger)
    # print(doc.ner.as_json)
    # print(doc.spans[:])
    result = []
    for span in doc.spans:
        result.append({'text': span.text, 'type': span.type})
        print(span.text, span.type)
    return jsonify(result)
