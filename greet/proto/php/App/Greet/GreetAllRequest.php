<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greet.proto

namespace App\Greet;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>greet.GreetAllRequest</code>
 */
class GreetAllRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.greet.Greeting greeting = 1;</code>
     */
    protected $greeting = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \App\Greet\Greeting $greeting
     * }
     */
    public function __construct($data = NULL) {
        \App\GPBMetadata\Greet::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.greet.Greeting greeting = 1;</code>
     * @return \App\Greet\Greeting|null
     */
    public function getGreeting()
    {
        return $this->greeting;
    }

    public function hasGreeting()
    {
        return isset($this->greeting);
    }

    public function clearGreeting()
    {
        unset($this->greeting);
    }

    /**
     * Generated from protobuf field <code>.greet.Greeting greeting = 1;</code>
     * @param \App\Greet\Greeting $var
     * @return $this
     */
    public function setGreeting($var)
    {
        GPBUtil::checkMessage($var, \App\Greet\Greeting::class);
        $this->greeting = $var;

        return $this;
    }

}

