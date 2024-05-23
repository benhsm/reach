## What and Why

Reach is a [bubbletea](https://github.com/charmbracelet/huh/tree/main/examples/bubbletea) TUI app that should probably have been a couple web forms connected to a spreadsheet. It was slapped together in a few days for a hackathon run by [HealthyGamerGG](https://www.healthygamer.gg/) for [mental health May](https://explore.healthygamer.gg/mayke-it).

The name is maybe a cheesy initialism for "Reflection on Experiences and Cognitions Helper", or something. More importantly, it is meant to evoke "reaching beyond your comfort-zone towards your goals", or "reaching out to other people".

The problem: anxiety and fear can narrow a person's world by inclining them to avoid contexts and situations that provoke those feelings. Avoidance behaviours can help someone keep safe, but become problematic when they don't accurately reflect danger, and prevent people from approaching their goals and/or functioning in ways they care about. What's difficult about this is that so long as someone continues to avoid something, they won't be able to get much evidence from experience that would demonstrate that they don't actually need to avoid it.

[Exposure therapy](https://en.wikipedia.org/wiki/Exposure_therapy) is a technique applied in clinical treatment of anxiety disorders which interrupts this cycle by repeatedly exposing patients to feared stimuli\*. By repeatedly encountering something they have irrationally feared, and noting again and again that there was no danger in those encounters, patients can learn a sense of safety in relation to those fears, and can become free from the compulsion to escape or avoid.

Reach is intended to function as an organizational aid for those interested in helping themselves overcome their anxieties and aversions on their own, using a similar pattern (and have a preference for keyboard driven, computer terminal interfaces).

\* (in case it needs to be said, this software is not aiming to be a substitute for any kind of professional mental health treatment.)

## Usage

The way it works:

1. Reach will present you with a form prompting you to specify something you would like to do, but may be afraid to for whatever reason. This form includes questions prompts for reflection on the thoughts and feelings that arise when you consider taking this action, and what concrete steps you would take to start on it, and will ask you to rate the task based on approximately how averse you are to it.
2. After planning out a few of these actions, the idea is that some point you'd actually get up, get out of the shell, and try to carry one out. When you come back, Reach provides you with another form to reflect on what you did, what happened, and whether or not it was actually as bad as you thought it would be.
3. All this collected information is stored in a database and presented in a tabular format, juxtaposing data on initial fears with recorded outcomes. The hope is that over time, the actions you plan and reflect on in this way could form a robust dataset that could help you calibrate your fears with reality, and perhaps overcome them as appropriate.

## Installation

Reach builds into a single binary. You can do:

```console
$ go install github.com/benhsm/reach 
```

Or install from source:

```
$ git clone https://github.com/benhsm/reach.git
$ cd reach
$ go build
$ ./reach
```

