import Inferno, { linkEvent } from "inferno";

const onDelete = (props, event) => {
  console.log("onDelete called");
  props.deleteVar(props.pair);
};

const Pair = props =>
  <div className="pair">
    <div onClick={linkEvent(props, onDelete)} className="action delete Red">
      rm
    </div>
    <div className="name Cyan">
      {props.pair.name}
    </div>
    <div id={props.pair.name} contenteditable className="value Green">
      {props.pair.value}
    </div>
  </div>;

export default Pair;
