import angular from 'angular';

angular.module('portainer.app').component('endpointStatusTag', {
  templateUrl: './endpoint-status-tag.html',
  bindings: {
    status: '<',
    endpointType: '<',
    edgeId: '<',
    emptyValue: '@',
  },
});
