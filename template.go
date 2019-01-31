package main

import "text/template"

var clientTemplate = template.Must(template.New("client").Parse(`<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: {{.SourceFileName}}

namespace {{ .Namespace }};

use Google\Protobuf\Internal\Message;

class {{ .ClassName }} implements {{ .InterfaceName }}
{
    /**
     * @var \GuzzleHttp\Client 
     */
    private $client;
    private $useJson;

    public function __construct(\GuzzleHttp\Client $client, $useJson = false)
    {
        $this->client = $client;
        $this->useJson = $useJson;
    }
{{ range .Methods }}
    public function {{ .PHPMethodName }}({{ .InputType }} ${{ .InputArg }}): {{ .OutputType }}
    {
        $res = $this->makeRequest('{{ .RPCMethodName }}', $this->serialize(${{ .InputArg }}));
        $out = new {{ .OutputType }}();
        $this->deserialize($out, $res->getBody()->getContents());
        return $out;
    }
{{ end }}
    private function makeRequest($method, $in)
    {
        $res = $this->client->post(self::SERVICE_NAME . '/' . $method, [
            'body' => $in,
            'http_errors' => false,
            'headers' => [
                'Content-Type' => $this->useJson ? 'application/json' : 'application/protobuf'
            ]
        ]);
        if ($res->getStatusCode() != 200) {
            throw new {{ .ExceptionName }}($res);
        }
        return $res;
    }

    private function serialize(Message $message)
    {
        if ($this->useJson) {
            return $message->serializeToJsonString();
        }
        return $message->serializeToString();
    }

    private function deserialize(Message $message, $data)
    {
        if ($this->useJson) {
            return $message->mergeFromJsonString($data);
        }
        return $message->mergeFromString($data);
    }
}
`))

var interfaceTemplate = template.Must(template.New("interface").Parse(`<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: {{.SourceFileName}}

namespace {{ .Namespace }};

interface {{ .InterfaceName }}
{
    const SERVICE_NAME = '{{ .ServiceName }}';
{{- with $svc := . }}
{{ range .Methods }}
    /**
     * @rpc {{ $svc.ServiceName }}/{{ .RPCMethodName }}
     * @param {{ .InputType }} ${{ .InputArg }}
     * @return {{ .OutputType }}
     */
    public function {{ .PHPMethodName }}({{ .InputType }} ${{ .InputArg }}): {{ .OutputType }};
{{ end }}
{{- end -}}
}
`))

var stubTemplate = template.Must(template.New("stub").Parse(`<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: {{.SourceFileName}}

namespace {{ .Namespace }};

class {{ .StubName }} implements {{ .InterfaceName }}
{
{{ range .Methods }}
    public ${{ .PHPCallbackName }};
{{- end }}
{{ range .Methods }}
    public function {{ .PHPMethodName }}({{ .InputType }} ${{ .InputArg }}): {{ .OutputType }}
    {
        if ($this->{{ .PHPCallbackName }}) {
            $func = $this->{{ .PHPCallbackName }};
            return $func(${{ .InputArg }});
        }
        throw new \BadMethodCallException("Method not stubbed");
    }
{{ end }}
}
`))

var noopTemplate = template.Must(template.New("noop").Parse(`<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: {{.SourceFileName}}

namespace {{ .Namespace }};

class {{ .NoopName }} implements {{ .InterfaceName }}
{
{{ range .Methods }}
    public function {{ .PHPMethodName }}({{ .InputType }} ${{ .InputArg }}): {{ .OutputType }}
    {
        return new {{ .OutputType }}();
    }
{{ end }}
}
`))

var exceptionTemplate = template.Must(template.New("exception").Parse(`<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: {{.SourceFileName}}

namespace {{ .Namespace }};

use Psr\Http\Message\ResponseInterface;

class {{ .ExceptionName }} extends \Exception
{
    private $twirpCode;
    private $meta;

    public function __construct(ResponseInterface $response)
    {
        $data = @json_decode($response->getBody()->getContents(), true);

        if (isset($data['msg'])) {
            $msg = $data['msg'];
        } else {
            $msg = $response->getStatusCode() . ': Unknown server response: ' . $response->getBody()->getContents();
        }

        $this->twirpCode = isset($data['code']) ? $data['code'] : 'internal';
        $this->meta = isset($data['meta']) ? $data['meta']: [];
        parent::__construct($msg, $response->getStatusCode());
    }

    public function getTwirpCode(): string
    {
        return $this->twirpCode;
    }

    public function getMeta(): array
    {
        return $this->meta;
    }
}
`))
