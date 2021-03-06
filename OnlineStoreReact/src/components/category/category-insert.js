import React, { Component } from "react";
import { Link } from "react-router-dom";
import { addCategory } from "../../actions/index";
import { bindActionCreators } from "redux";
import { connect } from "react-redux";
import { Field, reduxForm } from "redux-form";
import TextField from "material-ui/TextField";
import SelectField from "material-ui/SelectField";
import MenuItem from "material-ui/MenuItem";
import MuiThemeProvider from "material-ui/styles/MuiThemeProvider";
import getMuiTheme from "material-ui/styles/getMuiTheme";
import PropTypes from "prop-types"

const validate = values => {
    const errors = {}
    const requiredFields = ['Name', 'Description']
    requiredFields.forEach(field => {
        if (!values[field]) {
            errors[field] = 'Required'
        }
    })
    return errors
}

const renderTextField = ({ input, label, type, meta: { touched, error }, ...custom }) => (

    <TextField hintText={label}
        floatingLabelText={label}
        className="text-field"
        type={type}
        errorText={touched && error}
        {...custom}
        {...input}
    />
)

export class CategoryInsertForm extends Component {

    onSubmit(props) {
        console.log(props);
        this.props.addCategory(props).then(() => {
            this.context.router.history.push("/admin/categories");
            this.context.router.history.replace("/admin/categories");
        });
    }

    render() {

        const { handleSubmit, pristine, reset, submitting } = this.props
        return (
            <div className="col-md-12 col-sm-12">
                <div className="page-header">
                    <h4 className="text-primary"> Add a new Category </h4>
                </div>
                <div className="container">
                    <MuiThemeProvider muiTheme={getMuiTheme()}>
                        <form onSubmit={handleSubmit(this.onSubmit.bind(this))}>
                            <div className="form-group">
                                <Field type="text" name="Name" component={renderTextField} label="Category Name" />
                            </div>
                            <div>
                                <Field name="Description" component={renderTextField} label="Description" multiLine={true} rows={2} />
                            </div>
                            <div>
                                <button type="submit" className="btn btn-primary" disabled={pristine || submitting}>Save</button>
                                <button type="button" className="btn btn-default" disabled={pristine || submitting} onClick={reset}>Clear Values </button>
                                <Link to="/admin/categories" className="btn btn-default">Go Back</Link>
                            </div>
                        </form>
                    </MuiThemeProvider>
                </div>
            </div>
        )
    }
}

function mapDispatchToProps(dispatch) {
    return bindActionCreators({
        addCategory: addCategory,
    }, dispatch);
}


export default connect(null, mapDispatchToProps)(reduxForm({
    form: 'CategoryInsertForm',
    validate
})(CategoryInsertForm));

CategoryInsertForm.contextTypes = {
    router: PropTypes.object
}


// CategoryInsertForm = reduxForm({
//     form: 'CategoryInsertForm',
//     validate
// })(CategoryInsertForm)

// CategoryInsertForm = connect(null, mapDispatchToProps)(CategoryInsertForm)

// export default CategoryInsertForm