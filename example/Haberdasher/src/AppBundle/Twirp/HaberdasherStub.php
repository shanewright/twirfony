<?php
# Generated by protoc-gen-twirp_php, DO NOT EDIT.
# source: haberdasher.proto

namespace AppBundle\Twirp;

class HaberdasherStub implements HaberdasherInterface
{

    public $onMakeHat;

    public function makeHat(Size $size): Hat
    {
        if ($this->onMakeHat) {
            $func = $this->onMakeHat;
            return $func($size);
        }
        throw new \BadMethodCallException("Method not stubbed");
    }

}