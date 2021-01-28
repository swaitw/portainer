import angular from 'angular';

import controller from './endpoint-status-tag.controller';

angular.module('portainer.app').component('endpointStatusTag', {
  templateUrl: './endpoint-status-tag.html',
  controller,
  bindings: {
    status: '<',
    endpointType: '<',
    edgeId: '<',
    emptyValue: '@',
  },
});
