import Comments from './components/Comments';
import EthForm from './components/EthFormV2';

export default {
  routes: [
    {path: "/", name: "Comments", component: Comments},
    {path: "/form", name: "EthForm", component: EthForm},
  ]
}
