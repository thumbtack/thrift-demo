<?php
namespace tt\demo;

/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
use Thrift\Base\TBase;
use Thrift\Type\TType;
use Thrift\Type\TMessageType;
use Thrift\Exception\TException;
use Thrift\Exception\TProtocolException;
use Thrift\Protocol\TProtocol;
use Thrift\Protocol\TBinaryProtocolAccelerated;
use Thrift\Exception\TApplicationException;


final class Operation {
  const ADD = 0;
  const SUB = 1;
  const MUL = 2;
  const DIV = 3;
  static public $__names = array(
    0 => 'ADD',
    1 => 'SUB',
    2 => 'MUL',
    3 => 'DIV',
  );
}

class CalculateInput {
  static $_TSPEC;

  /**
   * @var double
   */
  public $number1 = null;
  /**
   * @var double
   */
  public $number2 = null;
  /**
   * @var int
   */
  public $op = null;

  public function __construct($vals=null) {
    if (!isset(self::$_TSPEC)) {
      self::$_TSPEC = array(
        1 => array(
          'var' => 'number1',
          'type' => TType::DOUBLE,
          ),
        2 => array(
          'var' => 'number2',
          'type' => TType::DOUBLE,
          ),
        3 => array(
          'var' => 'op',
          'type' => TType::I32,
          ),
        );
    }
    if (is_array($vals)) {
      if (isset($vals['number1'])) {
        $this->number1 = $vals['number1'];
      }
      if (isset($vals['number2'])) {
        $this->number2 = $vals['number2'];
      }
      if (isset($vals['op'])) {
        $this->op = $vals['op'];
      }
    }
  }

  public function getName() {
    return 'CalculateInput';
  }

  public function read($input)
  {
    $xfer = 0;
    $fname = null;
    $ftype = 0;
    $fid = 0;
    $xfer += $input->readStructBegin($fname);
    while (true)
    {
      $xfer += $input->readFieldBegin($fname, $ftype, $fid);
      if ($ftype == TType::STOP) {
        break;
      }
      switch ($fid)
      {
        case 1:
          if ($ftype == TType::DOUBLE) {
            $xfer += $input->readDouble($this->number1);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 2:
          if ($ftype == TType::DOUBLE) {
            $xfer += $input->readDouble($this->number2);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        case 3:
          if ($ftype == TType::I32) {
            $xfer += $input->readI32($this->op);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        default:
          $xfer += $input->skip($ftype);
          break;
      }
      $xfer += $input->readFieldEnd();
    }
    $xfer += $input->readStructEnd();
    return $xfer;
  }

  public function write($output) {
    $xfer = 0;
    $xfer += $output->writeStructBegin('CalculateInput');
    if ($this->number1 !== null) {
      $xfer += $output->writeFieldBegin('number1', TType::DOUBLE, 1);
      $xfer += $output->writeDouble($this->number1);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->number2 !== null) {
      $xfer += $output->writeFieldBegin('number2', TType::DOUBLE, 2);
      $xfer += $output->writeDouble($this->number2);
      $xfer += $output->writeFieldEnd();
    }
    if ($this->op !== null) {
      $xfer += $output->writeFieldBegin('op', TType::I32, 3);
      $xfer += $output->writeI32($this->op);
      $xfer += $output->writeFieldEnd();
    }
    $xfer += $output->writeFieldStop();
    $xfer += $output->writeStructEnd();
    return $xfer;
  }

}

class InputException extends TException {
  static $_TSPEC;

  /**
   * @var string
   */
  public $message = null;

  public function __construct($vals=null) {
    if (!isset(self::$_TSPEC)) {
      self::$_TSPEC = array(
        1 => array(
          'var' => 'message',
          'type' => TType::STRING,
          ),
        );
    }
    if (is_array($vals)) {
      if (isset($vals['message'])) {
        $this->message = $vals['message'];
      }
    }
  }

  public function getName() {
    return 'InputException';
  }

  public function read($input)
  {
    $xfer = 0;
    $fname = null;
    $ftype = 0;
    $fid = 0;
    $xfer += $input->readStructBegin($fname);
    while (true)
    {
      $xfer += $input->readFieldBegin($fname, $ftype, $fid);
      if ($ftype == TType::STOP) {
        break;
      }
      switch ($fid)
      {
        case 1:
          if ($ftype == TType::STRING) {
            $xfer += $input->readString($this->message);
          } else {
            $xfer += $input->skip($ftype);
          }
          break;
        default:
          $xfer += $input->skip($ftype);
          break;
      }
      $xfer += $input->readFieldEnd();
    }
    $xfer += $input->readStructEnd();
    return $xfer;
  }

  public function write($output) {
    $xfer = 0;
    $xfer += $output->writeStructBegin('InputException');
    if ($this->message !== null) {
      $xfer += $output->writeFieldBegin('message', TType::STRING, 1);
      $xfer += $output->writeString($this->message);
      $xfer += $output->writeFieldEnd();
    }
    $xfer += $output->writeFieldStop();
    $xfer += $output->writeStructEnd();
    return $xfer;
  }

}


