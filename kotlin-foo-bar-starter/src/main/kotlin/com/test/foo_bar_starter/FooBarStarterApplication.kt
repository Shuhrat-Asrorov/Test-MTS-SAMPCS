package com.test.foo_bar_starter

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class FooBarStarterApplication

fun main(args: Array<String>) {
	runApplication<FooBarStarterApplication>(*args)
}
