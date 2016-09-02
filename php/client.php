<?php

require_once(__DIR__ . '/vendor/autoload.php');
require_once(__DIR__ . '/../gen-php/tt/demo/Types.php');
require_once(__DIR__ . '/../gen-php/tt/demo/DemoService.php');

// Print use statement each time
echo "Use $argv[0] <host> <port>\n\n";

// defaults
$host = isset($argv[1]) ? $argv[1] : 'localhost';
$port = isset($argv[2]) ? intval($argv[2]) : 8080;

// actual Thrift boostrapping

// Create HTTP Transport
$transport = new \Thrift\Transport\TCurlClient(
    $host,
    $port,
    '/' // endpoint
);

// binary protocol
$protocol = new \Thrift\Protocol\TBinaryProtocol($transport);

// create a client
$client = new \tt\demo\DemoServiceClient($protocol);

// at this point, client is ready, and we can make RPC calls:
echo 'Random numbers: ' . PHP_EOL;
print_r($client->randomGenerator(1, 10, 5));

$input = new \tt\demo\CalculateInput();
$input->number1 = 5;
$input->number2 = 2;
$input->op = \tt\demo\Operation::ADD;

echo 'Calculate output: ' . $client->calculate($input) . PHP_EOL;
