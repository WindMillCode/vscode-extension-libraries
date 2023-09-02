// ignore_for_file: prefer_const_constructors, prefer_const_literals_to_create_immutables

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

class TemplateRiverpodProviderValue {

}

var TemplateRiverpodProviderInstance =
    TemplateRiverpodProviderValue();

class TemplateRiverpodNotifier  extends Notifier<TemplateRiverpodProviderValue>{


  @override
  TemplateRiverpodProviderValue build() {
    return TemplateRiverpodProviderInstance;
  }
}

final TemplateProvider = NotifierProvider<
    TemplateRiverpodNotifier,
    TemplateRiverpodProviderValue>(() {
  return TemplateRiverpodNotifier();
});

