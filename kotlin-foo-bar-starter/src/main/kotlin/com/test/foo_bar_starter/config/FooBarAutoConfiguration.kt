package com.test.foo_bar_starter.config

import com.test.foo_bar_starter.filter.FooBarFilter
import jakarta.servlet.Filter
import org.springframework.boot.autoconfigure.condition.ConditionalOnWebApplication
import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
@ConditionalOnWebApplication
class FooBarAutoConfiguration {

    @Bean
    fun fooBarFilter(): Filter {
        return FooBarFilter()
    }
}