import Inferno from "inferno";
import { connect } from "inferno-redux";
import { bindActionCreators } from "redux";
import EnvActions from "./EnvActions";
import Pair from "./Pair";

const Env = props => {
  let { deleteVar } = props;
  console.log(props);
  return (
    <div className="env">
      {props.env.map(e => <Pair deleteVar={deleteVar} pair={e} />)}
    </div>
  );
};
const mapStateToProps = (state, props) => ({ ...state, ...props });
const mapDispatchToProps = dispatch => ({
  ...bindActionCreators(EnvActions, dispatch)
});

export default connect(mapStateToProps, mapDispatchToProps)(Env);
