#!/bin/bash

bastet name="Gautam Dey" greeting="Hello" date="13 July" header.tpl body.tpl | diff --text  expected.output -

