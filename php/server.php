<?php
// everything uses fqdn for demo purposes

// run server using:
// php -S localhost:8080 server.php

// This script will end as soon as http request is finished


require_once(__DIR__ . '/vendor/autoload.php');
require_once(__DIR__ . '/../gen-php/tt/demo/Types.php');
require_once(__DIR__ . '/../gen-php/tt/demo/DemoService.php');

/**
 * @class DemoService
 *
 * Implements DemoService as defined in Thrift file.
 */
class DemoService implements \tt\demo\DemoServiceIf {
    /**
     * @param \tt\demo\CalculateInput $input
     * @return double
     * @throws \tt\demo\InputException
     */
    public function calculate(\tt\demo\CalculateInput $input) {
        switch ($input->op) {
            case \tt\demo\Operation::ADD:
                return $input->number1 + $input->number2;
            case \tt\demo\Operation::SUB:
                return $input->number1 - $input->number2;
            case \tt\demo\Operation::MUL:
                return $input->number1 * $input->number2;
            case \tt\demo\Operation::DIV:
                if ($input->number2 === .0) {
                    throw new \tt\demo\InputException([
                        'message' => 'Division by zero',
                    ]);
                }
                return $input->number1 - $input->number2;
        }
    }

    /**
     * @param int $min
     * @param int $max
     * @param int $numbers
     * @return int[]
     */
    public function randomGenerator($min, $max, $numbers) {
        $results = [];
        for ($i = 0; $i < $numbers; $i++) {
            // we don't care about security :)
            $results[] = rand($min, $max);
        }
        return $results;
    }
}

// initialize the service
$service = new DemoService();

// bootstrapping
$processor = new \tt\demo\DemoServiceProcessor($service);

// this will read php://stdin or php://input and print to php://stdout or php://output
$rw = \Thrift\Transport\TPhpStream::MODE_R | \Thrift\Transport\TPhpStream::MODE_W;
$transport = new \Thrift\Transport\TPhpStream($rw);
$transport->open();

// second and thrift argument set strict reads and writes
$protocol = new \Thrift\Protocol\TBinaryProtocol(
    $transport,
    true, // strict read - thrift lib will guarantee defined types, or will error
    true // strict write
);

// Actual logic happens here: processor will read from $protocol, which will read from
// $transport - which uses PHP Stream. After processor reads what needs to be done, it will execute
// the requested function and then wrihte to $protocol, which writes to $transport, which again uses
// PHP Stream.
$processor->process($protocol, $protocol);

$transport->close();
