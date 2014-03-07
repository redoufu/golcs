<?php

class shit{
	private $color;
	private $lenth;
	public function __construct($color='black', $lenth=10){
		$this->color = $color;
		$this->lenth = $lenth;
	}
	public function bomb(){
		echo "BOMB!\nCongratuations! You'v got a {$this->lenth}cm shit in {$this->color}.\n";
	}

}

$yourshit = new shit("yellow",25);
$yourshit->bomb();

?>
