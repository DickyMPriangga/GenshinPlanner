import React, { Component } from "react"
import axios from "axios"
import { Header, Input, List, Button, Image, Form, Divider } from "semantic-ui-react"

let endpoint = "http://localhost:8080"
const characterOpt = [
    { key: 'Albedo', value: 'Amber', text: 'Amber'},
    { key: 'Aloy', value: 'Aloy', text: 'Aloy'},
    { key: 'Amber', value: 'Amber', text: 'Amber'},
    { key: 'Barbara', value: 'Barbara', text: 'Barbara'},
    { key: 'Beidou', value: 'Beidou', text: 'Beidou'},
    { key: 'Bennett', value: 'Bennett', text: 'Bennett'},
    { key: 'Diluc', value: 'Diluc', text: 'Diluc'},
    { key: 'Diona', value: 'Diona', text: 'Diona'},
    { key: 'Fischl', value: 'Fischl', text: 'Fischl'},
    { key: 'Ganyu', value: 'Ganyu', text: 'Ganyu'},
    { key: 'Hu Tao', value: 'Hu Tao', text: 'Hu Tao'},
    { key: 'Jean', value: 'Jean', text: 'Jean'},
    { key: 'Kaeya', value: 'Kaeya', text: 'Kaeya'},
    { key: 'Kamisato Ayaka', value: 'Kamisato Ayaka', text: 'Kamisato Ayaka'},
    { key: 'Keqing', value: 'Keqing', text: 'Keqing'},
    { key: 'Klee', value: 'Klee', text: 'Klee'},
    { key: 'Kujou Sara', value: 'Kujou Sara', text: 'Kujou Sara'},
    { key: 'Lisa', value: 'Lisa', text: 'Lisa'},
    { key: 'Ningguang', value: 'Ningguang', text: 'Ningguang'},
    { key: 'Noelle', value: 'Noelle', text: 'Noelle'},
    { key: 'Qiqi', value: 'Qiqi', text: 'Qiqi'},
    { key: 'Raiden Shogun', value: 'Raiden Shogun', text: 'Raiden Shogun'},
    { key: 'Razor', value: 'Razor', text: 'Razor'},
    { key: 'Rosaria', value: 'Rosaria', text: 'Rosaria'},
    { key: 'Sucrose', value: 'Sucrose', text: 'Sucrose'},
    { key: 'Tartaglia', value: 'Tartaglia', text: 'Tartaglia'},
    { key: 'Traveler', value: 'Traveler', text: 'Traveler'},
    { key: 'Venti', value: 'Venti', text: 'Venti'},
    { key: 'Xiangling', value: 'Xiangling', text: 'Xiangling'},
    { key: 'Xiao', value: 'Xiao', text: 'Xiao'},
    { key: 'Xinyan', value: 'Xinyan', text: 'Xinyan'},
    { key: 'Yanfei', value: 'Yanfei', text: 'Yanfei'},
    { key: 'Yoimiya', value: 'Yoimiya', text: 'Yoimiya'},
    { key: 'Zhongli', value: 'Zhongli', text: 'Zhongli'},
]

class GenshinPlanner extends Component {
    constructor(props) {
        super(props)

        this.state = {
            charactername: "Amber",
            currentlevel: 1,
            targetlevel: 1,
            plans: []
        }
    }

    componentDidMount() {
        this.getPlan();
    }

    onChange = event => {
        this.setState({
            [event.target.name]: event.target.valueAsNumber
        });
    };

    onSelectChange = (event,data) => {
        this.setState({
            [data.name]: data.value
        })
    }

    onSubmit = () => {
        let plan = this.state;

        if(plan) {
            axios.post(
                endpoint + "/api/plan",
                {
                    "charactername" : this.state.charactername,
                    "currentlevel":this.state.currentlevel,
                    "targetlevel":this.state.targetlevel,
                },
                {
                    headers : {
                        "Content-Type": "application/x-www-form-urlencoded"
                    }
                })
                .then(res => {
                    this.getPlan();
                });
        }
    };

    onUpdate = (plan) => {
        if(plan) {
            axios.put(
                endpoint + "/api/plan",
                {
                    "_id" : plan._id,
                    "charactername" : plan.charactername,
                    "currentlevel" : plan.currentlevel,
                    "targetlevel" : plan.targetlevel,
                },
                {
                    headers : {
                        "Content-Type": "application/x-www-form-urlencoded"
                    }
                })
                .then(res => {
                    this.getPlan();
                });
        }
    }

    getPlan = () => {
        axios.get(endpoint + "/api/plan").then(res => {
            if (res.data) {
                this.setState({
                    plans: res.data.map(plan => {
                        var imgPath = '/charaImg/' + plan.charactername + '.png'
                        return (
                            <List.Item key={plan._id}>
                                <List divided size='large'>
                                    <List.Item>
                                        <Header as='h2'>
                                            <Image avatar src={imgPath}/>
                                            {plan.charactername}
                                        </Header>
                                        <h4>Required EXP</h4>
                                        <Input
                                            type="number"
                                            name="requiredEXP"
                                            value={plan.requiredexp}
                                        />
                                    </List.Item>
                                    <List.Item>
                                        <Form>
                                            <Form.Field
                                                type='number'
                                                label='Current LV'
                                                control='input'
                                                value={plan.currentlevel}
                                                inline
                                            />
                                            <Form.Field
                                                type='number'
                                                label='Target LV'
                                                control='input'
                                                value={plan.targetlevel}
                                                inline
                                            />
                                        </Form>
                                    </List.Item>
                                    <List.Item>
                                        <Button onClick={()=>this.deletePlan(plan._id)}>Delete</Button>
                                    </List.Item>
                                </List>
                            </List.Item>
                        );
                    })
                })
            }
        });
    };

    deletePlan = id => {
        axios
          .delete(endpoint + "/api/plan/" + id, {
            headers: {
              "Content-Type": "application/x-www-form-urlencoded"
            }
          })
          .then(res => {
            console.log(res);
            this.getPlan();
          });
      };

    render() {
        return (
            <div>
                <div className="row">
                    <Header className="header" as="h2">Genshin Impact Planner</Header>
                </div>
                <Divider />
                <div className="row">
                    <Form onSubmit={this.onSubmit}>
                        <Form.Group widths='equal'>
                            <Form.Select
                                placeholder='Select a Character'
                                name="charactername"
                                onChange={this.onSelectChange}
                                value={this.state.charactername}
                                options={characterOpt}
                                search
                            />
                            <Form.Field
                                type='number'
                                label='Current LV'
                                control='input'
                                name="currentlevel"
                                min='1'
                                max={this.state.targetlevel}
                                onChange={this.onChange}
                                value={this.state.currentlevel}
                                inline
                            />
                            <Form.Field
                                type='number'
                                label='Target LV'
                                control='input'
                                name="targetlevel"
                                min={this.state.currentlevel}
                                max='90'
                                onChange={this.onChange}
                                value={this.state.targetlevel}
                                inline
                            />
                        </Form.Group>
                        <Button type='submit'>Add</Button>
                    </Form>
                </div>
                <Divider />
                <div className="row">
                    <List divided verticalAlign="middle" size="big">
                        {this.state.plans}
                    </List>
                </div>
            </div>
        )
    }
}

export default GenshinPlanner