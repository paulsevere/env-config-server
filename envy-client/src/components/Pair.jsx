import Inferno, { linkEvent } from "inferno";

const onDelete = (props, event) => {
  console.log("onDelete called");
  props.deleteVar(props.pair);
};

const Pair = props =>
  <div style={{ display: "flex", flex: "0 0 auto" }}>
    <div className="card g--2">
      {props.pair.name}
    </div>
    <div id={props.pair.name} contenteditable className="card g--4">
      {props.pair.value}
    </div>

    <div onClick={linkEvent(props, onDelete)} className="close-icon-holder">
      <i className="material-icons md-48">close</i>
    </div>
    {(props.editing &&
      <div className="done-icon-holder">
        <i className="material-icons md-48">check</i>
      </div>) ||
      <div className="edit-icon-holder">
        <i className="material-icons md-48">mode_edit</i>
      </div>}
  </div>;

export default Pair;
