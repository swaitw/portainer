export default class EndpointStatusTagController {
  statusClass() {
    return this.status === 2 ? 'label-danger' : 'label-success';
  }
}
